  # selpg
  golang for selpg(CLI)
  
  [开发 Linux 命令行实用程序](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html)
  
  	一、usage
	
		$ selpg -s1 -e1 input_file
		该命令将把“input_file”的第 1 页写至标准输出（也就是屏幕），因为这里没有重定向或管道。
		$ selpg -s1 -e1 < input_file
		该命令与示例 1 所做的工作相同，但在本例中，selpg 读取标准输入，而标准输入已被 shell／内核重定向为来自“input_file”而不是显式命名的文件名参数。输入的第 1 页被写至屏幕。
		$ other_command | selpg -s10 -e20
		“other_command”的标准输出被 shell／内核重定向至 selpg 的标准输入。将第 10 页到第 20 页写至 selpg 的标准输出（屏幕）。
		$ selpg -s10 -e20 input_file >output_file
		selpg 将第 10 页到第 20 页写至标准输出；标准输出被 shell／内核重定向至“output_file”。
		$ selpg -s10 -e20 input_file 2>error_file
		selpg 将第 10 页到第 20 页写至标准输出（屏幕）；所有的错误消息被 shell／内核重定向至“error_file”。请注意：在“2”和“>”之间不能有空格；这是 shell 语法的一部分（请参阅“man bash”或“man sh”）。
		$ selpg -s10 -e20 input_file >output_file 2>error_file
		selpg 将第 10 页到第 20 页写至标准输出，标准输出被重定向至“output_file”；selpg 写至标准错误的所有内容都被重定向至“error_file”。当“input_file”很大时可使用这种调用；您不会想坐在那里等着 selpg 完成工作，并且您希望对输出和错误都进行保存。
		$ selpg -s10 -e20 input_file >output_file 2>/dev/null
		selpg 将第 10 页到第 20 页写至标准输出，标准输出被重定向至“output_file”；selpg 写至标准错误的所有内容都被重定向至 /dev/null（空设备），这意味着错误消息被丢弃了。设备文件 /dev/null 废弃所有写至它的输出，当从该设备文件读取时，会立即返回 EOF。
		$ selpg -s10 -e20 input_file >/dev/null
		selpg 将第 10 页到第 20 页写至标准输出，标准输出被丢弃；错误消息在屏幕出现。这可作为测试 selpg 的用途，此时您也许只想（对一些测试情况）检查错误消息，而不想看到正常输出。
		$ selpg -s10 -e20 input_file | other_command
		selpg 的标准输出透明地被 shell／内核重定向，成为“other_command”的标准输入，第 10 页到第 20 页被写至该标准输入。“other_command”的示例可以是 lp，它使输出在系统缺省打印机上打印。“other_command”的示例也可以 wc，它会显示选定范围的页中包含的行数、字数和字符数。“other_command”可以是任何其它能从其标准输入读取的命令。错误消息仍在屏幕显示。
		$ selpg -s10 -e20 input_file 2>error_file | other_command
		与上面的示例 9 相似，只有一点不同：错误消息被写至“error_file”。

		二、测试（page_len = 20)

		1、./selpg -s1 -e1 -l10 1.txt

		line1
		line2
		line3
		line4
		line5
		line6
		line7
		line8
		line9
		line10
		
		2、./selpg -s1 -e1 1.txt >3.txt
		3.txt:
		line1
		line2
		line3
		line4
		line5
		line6
		line7
		line8
		line9
		line10
		line11
		line12
		line13
		line14
		line15
		line16
		line17
		line18
		line19
		line20
		
		3、./selpg -s1 -e4 1.txt 2>error.txt
		error.txt:
		DEBUG: before handling 1st arg
		DEBUG: before handling 2nd arg
		DEBUG: before while loop for opt args
		DEBUG: before check for filename arg
		DEBUG: argno = 3
		DEBUG: sa.start_page = 1
		DEBUG: sa.end_page = 4
		DEBUG: sa.page_len = 20
		DEBUG: sa.page_type = l
		DEBUG: sa.print_dest = 
		DEBUG: sa.in_filename = 1.txt
		end page is greater than total page
		
		4、./selpg -s1 -e1 -f 1.txt
		line1
		line2
		line3
		line4
		line5
		line6
		line7
		line8
		line9
		line10
		line11
		line12
		line13
		line14
		line15
		line16
		line17
		line18
		line19
		line20
		line21
		line22
		line23
		line24
		line25
		line26
		line27
		line28
		line29
		line30
		
		５、./selpg -s1 -e1 -dlp1 1.txt
		（通过cmd := exec.Command("cat","-n")也就是cat命令将1.txt的内容打印在屏幕上
		line1
		line2
		line3
		line4
		line5
		line6
		line7
		line8
		line9
		line10
		line11
		line12
		line13
		line14
		line15
		line16
		line17
		line18
		line19
		line20

		
		
