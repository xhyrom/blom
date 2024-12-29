.text
.globl main
main:
	pushq %rbp
	movq %rsp, %rbp
	movl $7, %edi
	callq lol
	movl %eax, %edi
	callq lol
	movl %eax, %edi
	callq lol
	movl %eax, %edi
	callq lol
	movl %eax, %edi
	callq lol
	movl %eax, %edi
	callq lol
	movl %eax, %edi
	callq lol
	movl %eax, %edi
	callq lol
	movl %eax, %edi
	callq lol
	leave
	ret
.type main, @function
.size main, .-main
/* end function main */

.text
lol:
	pushq %rbp
	movq %rsp, %rbp
	movl %edi, %eax
	leave
	ret
.type lol, @function
.size lol, .-lol
/* end function lol */

