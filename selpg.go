package main

import (
    "os"
    "fmt"
    "strconv"
    "bufio"
    "io"
    "os/exec"
)

const BUFSIZ = 16*1024
const MAX_IN = 18446
type selpg_args struct {
    start_page int
    end_page int
    in_filename string
    page_len int
    page_type int
    print_dest string
}

var ac int
var sp_args selpg_args
var sa selpg_args
var str [BUFSIZ]byte

var progname string

func usage() {
    fmt.Fprintf(os.Stderr,
    "\nUSAGE: %s -sstart_page -eend_page [ -f | -llines_per_page ][ -ddest ] [ in_filename ]\n", progname);
}

func process_args(av []string) {
    //var s1 [BUFSIZ]byte
    //var s2 [BUFSIZ]byte
    var argno int
    var i int
    if len(av) < 3 {
        fmt.Fprintf(os.Stderr, "%s: not enough arguments\n", progname)
        usage()
        os.Exit(1)
    }
    fmt.Fprintf(os.Stderr, "DEBUG: before handling 1st arg\n");
    //处理第一个参数
    if av[1][0] != '-'|| av[1][1] != 's' {
        fmt.Fprintf(os.Stderr, "%s: 1st arg should be -sstart_page\n", progname);
        usage()
        os.Exit(1)
    }
    i , _ = strconv.Atoi(av[1][2:])
    if i <1|| i > MAX_IN {
        fmt.Fprintf(os.Stderr, "%s: invalid start page %d\n", progname, i);
        usage()
        os.Exit(1)
    }
    sa.start_page = i

    fmt.Fprintf(os.Stderr, "DEBUG: before handling 2nd arg\n");

    if av[2][0] != '-'|| av[2][1] != 'e' {
        fmt.Fprintf(os.Stderr, "%s: 2nd arg should be -eend_page\n", progname);
        usage()
        os.Exit(1)
    }
    i ,_ = strconv.Atoi(av[2][2:])
    if i < 1||i > MAX_IN||i < sa.start_page {
        fmt.Fprintf(os.Stderr, "%s: invalid end page %d\n", progname, i);
        usage()
        os.Exit(1)
    }
    sa.end_page = i;

    fmt.Fprintf(os.Stderr, "DEBUG: before while loop for opt args\n");

    argno = 3;
    for {
        if argno > (ac-1) ||av[argno][0] != '-' {
            break
        }
        switch av[argno][1] {
            case 'l':
                i,_ = strconv.Atoi(av[argno][2:])
                if i < 1|| i > MAX_IN {
                    fmt.Fprintf(os.Stderr, "%s: invalid page length %d\n", progname, i);
                    usage()
                    os.Exit(1)
                }
                sa.page_len = i
                argno = argno+1
                continue
                break
            case 'f':
                if len(av[argno]) > 2 {
                    fmt.Fprintf(os.Stderr, "%s: option should be \"-f\"\n", progname);
                    usage()
                    os.Exit(1)
                }
                sa.page_type = 'f'
                argno = argno+1
                continue
                break
            case 'd':
                if len(av[argno]) <= 2 {
                    fmt.Fprintf(os.Stderr,
                    "%s: -d option requires a printer destination\n", progname);
                    usage()
                    os.Exit(1)
                }
                sa.print_dest = av[argno][2:]
                argno = argno+1
                continue
                break
            default:
                fmt.Fprintf(os.Stderr, "%s: unknown option", progname);
                usage()
                os.Exit(1)
                break
        }
    }

    fmt.Fprintf(os.Stderr, "DEBUG: before check for filename arg\n");
    fmt.Fprintf(os.Stderr, "DEBUG: argno = %d\n", argno);

    if argno <= (ac-1) {
        sa.in_filename = av[argno]
        infile, err := os.Open(sa.in_filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "%s: input file \"%s\" does not exist\n",
            progname, sa.in_filename);
            os.Exit(1)
        }
        defer infile.Close()
    }

    fmt.Fprintf(os.Stderr, "DEBUG: sa.start_page = %d\n", sa.start_page)
    fmt.Fprintf(os.Stderr, "DEBUG: sa.end_page = %d\n", sa.end_page)
    fmt.Fprintf(os.Stderr, "DEBUG: sa.page_len = %d\n", sa.page_len)
    fmt.Fprintf(os.Stderr, "DEBUG: sa.page_type = %c\n", sa.page_type)
    fmt.Fprintf(os.Stderr, "DEBUG: sa.print_dest = %s\n", sa.print_dest)
    fmt.Fprintf(os.Stderr, "DEBUG: sa.in_filename = %s\n", sa.in_filename)
}

func process_input() {
    var line_ctr int = 0
    var page_ctr int = 1
    
    if sa.print_dest != "" {
        f, err := os.Open(sa.in_filename)  
        if err != nil {
            fmt.Fprintf(os.Stderr, "%s: could not open input file \"%s\"\n",
            progname, sa.in_filename);
            os.Exit(1)  
        }
        cmd := exec.Command("cat","-n")
        stdin, err := cmd.StdinPipe()
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        buff := bufio.NewReader(f)
        for {  
            line, err := buff.ReadString('\f') //以'\n'为结束符读入一行
            if err != nil || io.EOF == err {
                break
            }
            line_ctr++
            if line_ctr > sa.page_len {
                page_ctr++
                line_ctr = 1
            }
            if page_ctr >= sa.start_page&& page_ctr <= sa.end_page {
                stdin.Write([]byte(line + "\n"))
            }
        }
        stdin.Close()
        cmd.Stdout = os.Stdout
        cmd.Run() 
    }

    if sa.in_filename != "" {
        f, err := os.Open(sa.in_filename)  
        if err != nil {
            fmt.Fprintf(os.Stderr, "%s: could not open input file \"%s\"\n",
            progname, sa.in_filename); 
        }
        defer f.Close()
        buff := bufio.NewReader(f) 
        if sa.page_type == 'l' {
            for {  
                line, err := buff.ReadString('\n') //以'\n'为结束符读入一行
                if err != nil || io.EOF == err {
                    break
                }
                line_ctr++
                if line_ctr > sa.page_len {
                    page_ctr++
                    line_ctr = 1
                }
                if page_ctr >= sa.start_page&& page_ctr <= sa.end_page {
                    fmt.Print(line)
                }
            }   
        } else {
            page_ctr = 1
            for {  
                line, err := buff.ReadString('\n') //以'\n'为结束符读入一行
                if err != nil || io.EOF == err {
                    break
                }
                if line == "\f" {
                    page_ctr++
                }
                if page_ctr >= sa.start_page&& page_ctr <= sa.end_page {
                    fmt.Print(line)
                }
            }   
        }
    }

    if page_ctr < sa.start_page {
        fmt.Fprintf(os.Stderr, "start page is greater than total page\n")
    } else if page_ctr < sa.end_page {
        fmt.Fprintf(os.Stderr,"end page is greater than total page\n") 
    }
}




func main() {
    var k int = 0
    var t int = 0
    var temp [BUFSIZ]byte
    args := os.Args //获取用户输入的所有参数
    ac = len(args)
    for j := len(args[0])-1; j >= 0; j-- {
        if args[0][j] == '/' {
            break;
        }
        temp[k] = args[0][j]
        k++
    }
    for j := k-1; j >= 0; j--{
        str[t] = temp[j]
        t++
    }
    progname = string(str[:])
    sa.start_page = 1
    sa.end_page = 1
    sa.in_filename = ""
    sa.page_len = 20
    sa.page_type = 'l'
    sa.print_dest = ""
    process_args(args)
    process_input()
}
