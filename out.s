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
	movl $9, %eax
	leave
	ret
.type main, @function
.size main, .-main
/* end function main */

.section .note.GNU-stack,"",@progbits
