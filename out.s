.data
.balign 8
main.0:
	.ascii "here\n"
	.byte 0
/* end data */

.data
.balign 8
main.1:
	.ascii "outer: %d\n"
	.byte 0
/* end data */

.text
.globl main
main:
	pushq %rbp
	movq %rsp, %rbp
	cmpl $11, %edi
	subq $16, %rsp
	movq %rsp, %rax
	movl $1244, (%rax)
	leaq main.0(%rip), %rdi
	movl $0, %eax
	callq printf
	movl $4294967295, %eax
	movq %rbp, %rsp
	subq $0, %rsp
	leave
	ret
.type main, @function
.size main, .-main
/* end function main */

.section .note.GNU-stack,"",@progbits
