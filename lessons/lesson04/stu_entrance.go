
scanner := bufio.NewScanner(os.Stdin)
for {
	var cmd string
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscan(line, &cmd)
	switch cmd {
	case "add":
		add()
	default:
		usage()
}
