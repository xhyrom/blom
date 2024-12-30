.text
.globl fib
fib:
	pushq %rbp
	movq %rsp, %rbp
	subq $8, %rsp
	pushq %rbx
	movl %edi, %ebx
	cmpl $1, %ebx
	jle .Lbb2
	movl %ebx, %edi
	subl $1, %edi
	callq fib
	xchgl %eax, %ebx
	movl %eax, %edi
	subl $2, %edi
	callq fib
	addl %ebx, %eax
	jmp .Lbb3
.Lbb2:
	movl %ebx, %eax
.Lbb3:
	popq %rbx
	leave
	ret
.type fib, @function
.size fib, .-fib
/* end function fib */

.text
.globl main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $16, %rsp
	movq %rsp, %rax
	movl $0, (%rax)
	subq $16, %rsp
	movq %rsp, %rsi
	movl $5, (%rsi)
	movl $5, %edx
	movl $0, %ecx
.Lbb7:
	cmpl $10, %edx
	jge .Lbb9
	addl %edx, %ecx
	movl %ecx, (%rax)
	addl $1, %edx
	movl %edx, (%rsi)
	jmp .Lbb7
.Lbb9:
	subq $16, %rsp
	movq %rsp, %rsi
	movl $0, (%rsi)
	movl $0, %edx
.Lbb11:
	cmpl $10, %edx
	jl .Lbb16
.Lbb12:
	cmpl $10, %ecx
	jge .Lbb14
	addl $1, %ecx
	movl %ecx, (%rax)
	jmp .Lbb12
.Lbb14:
	movl %ecx, %eax
	movq %rbp, %rsp
	subq $0, %rsp
	leave
	ret
.Lbb16:
	addl %edx, %ecx
	movl %ecx, (%rax)
	addl $1, %edx
	movl %edx, (%rsi)
	jmp .Lbb11
.type main, @function
.size main, .-main
/* end function main */

.section .note.GNU-stack,"",@progbits
