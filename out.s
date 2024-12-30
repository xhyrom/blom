.text
.globl hoh
hoh:
	pushq %rbp
	movq %rsp, %rbp
	movl $5, %eax
	leave
	ret
.type hoh, @function
.size hoh, .-hoh
/* end function hoh */

.text
.globl main
main:
	pushq %rbp
	movq %rsp, %rbp
	callq hoh
	leave
	ret
.type main, @function
.size main, .-main
/* end function main */

.section .note.GNU-stack,"",@progbits
