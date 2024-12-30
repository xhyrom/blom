.data
.balign 8
main.35:
	.ascii "Hello, World! %d | %d\n"
	.byte 0
/* end data */

.text
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
	subq $8, %rsp
	pushq %rbx
	subq $16, %rsp
	movq %rsp, %rax
	movl $0, (%rax)
	subq $16, %rsp
	movq %rsp, %rdx
	movl $5, (%rdx)
	movl $5, %ecx
	movl $0, %ebx
.Lbb7:
	cmpl $10, %ecx
	jge .Lbb9
	addl %ecx, %ebx
	movl %ebx, (%rax)
	addl $1, %ecx
	movl %ecx, (%rdx)
	jmp .Lbb7
.Lbb9:
	subq $16, %rsp
	movq %rsp, %rdx
	movl $0, (%rdx)
	movl $0, %ecx
.Lbb11:
	cmpl $10, %ecx
	jl .Lbb16
.Lbb12:
	cmpl $10, %ebx
	jge .Lbb14
	addl $1, %ebx
	movl %ebx, (%rax)
	jmp .Lbb12
.Lbb14:
	movl $40, %edi
	callq fib
	movl %eax, %esi
	movl %ebx, %eax
	movl $99, %edx
	leaq main.35(%rip), %rdi
	movl %eax, %ebx
	movl $0, %eax
	callq printf
	movl %ebx, %eax
	movq %rbp, %rsp
	subq $16, %rsp
	popq %rbx
	leave
	ret
.Lbb16:
	addl %ecx, %ebx
	movl %ebx, (%rax)
	addl $1, %ecx
	movl %ecx, (%rdx)
	jmp .Lbb11
.type main, @function
.size main, .-main
/* end function main */

.section .note.GNU-stack,"",@progbits
