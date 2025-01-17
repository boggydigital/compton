package compton

func SetTint(element Element, color string) {
	element.SetAttribute(
		"style",
		"--tint-bg:color-mix(in display-p3,"+color+" var(--cma),var(--c-background))")

}
