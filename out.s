.data
.balign 8
main.0:
	.ascii "fib(50) = %d\n"
	.byte 0
/* end data */

.text
fib:
	pushq %rbp
	movq %rsp, %rbp
	subq $8, %rsp
	pushq %rbx
	movq %rdi, %rbx
	cmpq $1, %rbx
	jle .Lbb2
	movq %rbx, %rdi
	subq $1, %rdi
	callq fib
	xchgq %rax, %rbx
	movq %rax, %rdi
	subq $2, %rdi
	callq fib
	addq %rbx, %rax
	jmp .Lbb3
.Lbb2:
	movq %rbx, %rax
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
	movl $40, %edi
	callq fib
	movq %rax, %rsi
	leaq main.0(%rip), %rdi
	movl $0, %eax
	callq printf
	movl $0, %eax
	leave
	ret
.type main, @function
.size main, .-main
/* end function main */

.section .note.GNU-stack,"",@progbits
