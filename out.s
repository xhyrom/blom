.data
.balign 8
.0:
	.ascii "here\n"
	.byte 0
/* end data */

.data
.balign 8
.1:
	.ascii "outer: %d\n"
	.byte 0
/* end data */

.text
.globl main
main:
	pushq %rbp
	movq %rsp, %rbp
	cmpl $11, %edi
	jg .Lbb4
	subq $16, %rsp
	movq %rsp, %rax
	movl $1244, (%rax)
	movl $1255, %esi
	jmp .Lbb6
.Lbb4:
	cmpl $11, %edi
	movl $2, %esi
.Lbb6:
	cmpl $2, %esi
	jz .Lbb8
	leaq .1(%rip), %rdi
	movl $0, %eax
	callq printf
	movl $0, %eax
	jmp .Lbb9
.Lbb8:
	leaq .0(%rip), %rdi
	movl $0, %eax
	callq printf
	movl $4294967295, %eax
.Lbb9:
	movq %rbp, %rsp
	subq $0, %rsp
	leave
	ret
.type main, @function
.size main, .-main
/* end function main */

.section .note.GNU-stack,"",@progbits
