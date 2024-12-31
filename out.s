.data
.balign 8
main.8:
	.ascii "inner a: %d\n"
	.byte 0
/* end data */

.data
.balign 8
main.10:
	.ascii "outer a: %d\n"
	.byte 0
/* end data */

.data
.balign 8
main.12:
	.ascii "a: %d\n"
	.byte 0
/* end data */

.text
.globl main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $8, %rsp
	pushq %rbx
	subq $16, %rsp
	movq %rsp, %rbx
	movl $7, (%rbx)
	movl $8, (%rbx)
	movl $8, %esi
	leaq main.8(%rip), %rdi
	movl $0, %eax
	callq printf
	movl $8, %esi
	leaq main.10(%rip), %rdi
	movl $0, %eax
	callq printf
	movl $9, (%rbx)
	movl $18, %esi
	leaq main.12(%rip), %rdi
	movl $0, %eax
	callq printf
	movl $18, %eax
	movq %rbp, %rsp
	subq $16, %rsp
	popq %rbx
	leave
	ret
.type main, @function
.size main, .-main
/* end function main */

.section .note.GNU-stack,"",@progbits
