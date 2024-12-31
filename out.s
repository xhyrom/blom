.data
.balign 8
main.8:
	.ascii "%d\n"
	.byte 0
/* end data */

.data
.balign 8
main.10:
	.ascii "%d\n"
	.byte 0
/* end data */

.data
.balign 8
main.12:
	.ascii "%d\n"
	.byte 0
/* end data */

.text
.globl main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $16, %rsp
	movq %rsp, %rax
	movl $7, (%rax)
	movl $9, (%rax)
	movl $9, %esi
	leaq main.8(%rip), %rdi
	movl $0, %eax
	callq printf
	movl $9, %esi
	leaq main.10(%rip), %rdi
	movl $0, %eax
	callq printf
	movl $5, %esi
	leaq main.12(%rip), %rdi
	movl $0, %eax
	callq printf
	movl $5, %eax
	movq %rbp, %rsp
	subq $0, %rsp
	leave
	ret
.type main, @function
.size main, .-main
/* end function main */

.section .note.GNU-stack,"",@progbits
