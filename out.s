.data
.balign 8
.0:
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
	subq $16, %rsp
	movq %rsp, %rcx
	addl %eax, %ebx
	movl %ebx, (%rcx)
	subq $16, %rsp
	movq %rsp, %rcx
	addl $1, %eax
	movl %eax, (%rcx)
	jmp .Lbb2
.Lbb4:
	movl %ebx, %eax
	movl %eax, %ebx
	subq $16, %rsp
	movq %rsp, %rax
	movl $0, (%rax)
	movl $0, %eax
.Lbb7:
	cmpl $10, %eax
	jl .Lbb13
.Lbb8:
	cmpl $10, %ebx
	jge .Lbb10
	subq $16, %rsp
	movq %rsp, %rax
	addl $1, %ebx
	movl %ebx, (%rax)
	jmp .Lbb8
.Lbb10:
	movl %ebx, %eax
	movl %eax, %esi
	leaq .0(%rip), %rdi
	movl %eax, %ebx
	movl $0, %eax
	callq printf
	movl %ebx, %eax
	movq %rbp, %rsp
	subq $16, %rsp
	popq %rbx
	leave
	ret
.Lbb13:
	subq $16, %rsp
	movq %rsp, %rcx
	addl %eax, %ebx
	movl %ebx, (%rcx)
	subq $16, %rsp
	movq %rsp, %rcx
	addl $1, %eax
	movl %eax, (%rcx)
	jmp .Lbb7
.type main, @function
.size main, .-main
/* end function main */

.section .note.GNU-stack,"",@progbits
