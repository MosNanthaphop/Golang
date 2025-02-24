// bank_test.go
package bank

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// ทดสอบการสร้างบัญชีออมทรัพย์ (savingsAccount) และการใช้งานเบื้องต้น
func TestSavingsAccount(t *testing.T) {
	acc := NewSavingsAccount(1000)

	// ทดสอบ GetBalance() หลังสร้าง
	if acc.GetBalance() != 1000 {
		t.Errorf("expected balance = 1000, but got %v", acc.GetBalance())
	}

	// ทดสอบฝากเงิน (Deposit)
	err := acc.Deposit(500)
	if err != nil {
		t.Errorf("unexpected error on deposit: %v", err)
	}
	if acc.GetBalance() != 1500 {
		t.Errorf("expected balance = 1500 after deposit, but got %v", acc.GetBalance())
	}

	// ทดสอบฝากเงินไม่ถูกต้อง (Deposit <= 0)
	err = acc.Deposit(0)
	if err == nil {
		t.Error("expected error when depositing 0, but got nil")
	}

	// ทดสอบถอนเงิน (Withdraw)
	err = acc.Withdraw(1000)
	if err != nil {
		t.Errorf("unexpected error on withdraw: %v", err)
	}
	if acc.GetBalance() != 500 {
		t.Errorf("expected balance = 500 after withdraw, but got %v", acc.GetBalance())
	}

	// ทดสอบถอนเงินเกินยอด (InsufficientFundsError)
	err = acc.Withdraw(600)
	if err == nil {
		t.Error("expected error on insufficient funds, but got nil")
	} else {
		// เช็คว่าเป็นประเภท InsufficientFundsError หรือไม่
		if _, ok := err.(*InsufficientFundsError); !ok {
			t.Errorf("expected InsufficientFundsError, but got %T", err)
		}
	}
}

// ทดสอบการสร้างบัญชีดอกเบี้ย (interestAccount)
func TestInterestAccount(t *testing.T) {
	intAcc := NewInterestAccount(2000, 0.05)

	// ทดสอบค่าเริ่มต้น
	if intAcc.GetBalance() != 2000 {
		t.Errorf("expected balance = 2000, but got %v", intAcc.GetBalance())
	}

	// ทดสอบ AddInterest()
	intAcc.AddInterest()
	// ดอกเบี้ย 5% ของ 2000 = 100
	if intAcc.GetBalance() != 2100 {
		t.Errorf("expected balance = 2100 after interest, but got %v", intAcc.GetBalance())
	}

	// ทดสอบ Withdraw (สังเกตว่าจะมีข้อความ "Withdrawing from interest account...")
	err := intAcc.Withdraw(100)
	if err != nil {
		t.Errorf("unexpected error on withdraw from interest account: %v", err)
	}
	if intAcc.GetBalance() != 2000 {
		t.Errorf("expected balance = 2000 after withdraw, but got %v", intAcc.GetBalance())
	}
}

// ทดสอบฟังก์ชันโอนเงิน (Transfer)
func TestTransfer(t *testing.T) {
	from := NewSavingsAccount(1000)
	to := NewSavingsAccount(500)

	// โอน 300 จาก from -> to
	err := Transfer(from, to, 300)
	if err != nil {
		t.Errorf("unexpected error on transfer: %v", err)
	}
	if from.GetBalance() != 700 {
		t.Errorf("expected from-balance = 700, but got %v", from.GetBalance())
	}
	if to.GetBalance() != 800 {
		t.Errorf("expected to-balance = 800, but got %v", to.GetBalance())
	}

	// ทดสอบโอนเกินยอด (ให้เกิด InsufficientFundsError)
	err = Transfer(from, to, 2000) // from มีแค่ 700
	if err == nil {
		t.Error("expected error on transfer exceeding balance, but got nil")
	} else {
		if _, ok := err.(*InsufficientFundsError); !ok {
			t.Errorf("expected InsufficientFundsError, but got %T", err)
		}
	}
}

// ทดสอบ RiskyBankOperation (Panic & Recover)
// วิธีทดสอบพฤติกรรม panic ใน Unit Test
func TestRiskyBankOperation(t *testing.T) {
	// 1) กรณีที่ไม่เกิด Panic
	// (เราจะเรียกฟังก์ชันให้ทำงาน ถ้าไม่มี panic, ทดสอบผ่าน)
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("did not expect a panic, but got one: %v", r)
		}
	}()
	RiskyBankOperation(100)
}

// ทดสอบกรณีเกิด Panic จริง ๆ
// ทดสอบกรณีฟังก์ชัน Recover ภายในตัวเอง (panic ไม่หลุดมาถึง test)
func TestRiskyBankOperationPanic(t *testing.T) {
	// 1) สำรอง stdout ปัจจุบันไว้ เพื่อจะได้คืนค่าทีหลัง
	originalStdout := os.Stdout

	// 2) สร้าง Pipe เพื่อดักจับ output
	r, w, _ := os.Pipe()
	os.Stdout = w // เปลี่ยนให้ stdout ชี้ไปที่ตัว Pipe

	// 3) เรียกฟังก์ชันที่มีการ recover ภายใน
	//    โดยใส่ค่าเป็นลบ (ซึ่งจะ panic แล้ว recover ในตัวเอง)
	RiskyBankOperation(-50)

	// 4) ปิด writer และคืนค่า stdout กลับ
	w.Close()
	os.Stdout = originalStdout

	// 5) อ่านข้อมูลจากตัว Pipe (r)
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// 6) ตรวจสอบข้อความที่คาดว่าจะเกิดขึ้น
	if !strings.Contains(output, "เกิด Panic แล้วทำการ Recover สำเร็จ:") {
		t.Errorf("Expected output to contain \"เกิด Panic แล้วทำการ Recover สำเร็จ:\", but got:\n%s", output)
	}

	// (Optional) ตรวจสอบว่าไม่หลุด crash มาถึง Test
	// ถ้ามี panic หลุดมาถึง Test จริง ๆ Test จะไม่มาถึงบรรทัดนี้ เพราะจะ Fail ก่อน
}
