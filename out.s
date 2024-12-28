.data
.balign 8
main.0:
	.ascii "i: %d\n"
	.byte 0
/* end data */

.text
.globl main
main:
	pushq %rbp
	movq %rsp, %rbp
	subq $8, %rsp
	pushq %rbx
	movl $0, %ebx
	movl $5, %eax
.Lbb2:
	cmpl $10, %eax
	jge .Lbb4
	addl %eax, %ebx
	addl $1, %eax
	jmp .Lbb2
.Lbb4:
	subq $16, %rsp
	movq %rsp, %rcx
	movl $0, (%rcx)
	movl $0, %eax
.Lbb6:
	cmpl $10, %eax
	jl .Lbb12
.Lbb7:
	cmpl $10, %ebx
	jge .Lbb9
	addl $1, %ebx
	jmp .Lbb7
.Lbb9:
	movl %ebx, %eax
	movl %eax, %esi
	leaq main.0(%rip), %rdi
	movl %eax, %ebx
	movl $0, %eax
	callq printf
	movl %ebx, %eax
	movq %rbp, %rsp
	subq $16, %rsp
	popq %rbx
	leave
	ret
.Lbb12:
	addl %eax, %ebx
	addl $1, %eax
	movl %eax, (%rcx)
	jmp .Lbb6
.type main, @function
.size main, .-main
/* end function main */

.section .note.GNU-stack,"",@progbits
