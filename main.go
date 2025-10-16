package main

import (
    "context"
    "fmt"
    "log"
    "google.golang.org/genai"
	"bufio"
	"strings"
	"os"
)

func main() {
	fmt.Println("||------------Welcome To Tanners Coding Project Prompter------------||")
    // Opening Menu:
	reader := bufio.NewReader(os.Stdin)
	var p string
	for {
		fmt.Println("Pick A Programming Language!")
		fmt.Println("||--------------------||")
		fmt.Println("1. Java")
		fmt.Println("2. Javascript")
		fmt.Println("3. Go/Golang")
		fmt.Println("4. Python")
		fmt.Println("5. C++")
		fmt.Println("6. C#")
		fmt.Println("7. quit")
		fmt.Println("||--------------------||")
		fmt.Println("Input a number: ")
		menuInput, _ := reader.ReadString('\n')
		p = strings.TrimSpace(menuInput)		

		if p == "1" {
			difficultyJava()
		} else if p == "2" {
			difficultyJavascript()
		} else if p == "3" {
			difficultyGo()
		} else if p == "4" {
			difficultyPython()
		} else if p == "5" {
			difficultyCPlus()
		} else if p == "6" {
			difficultyCSharp()
		} else if p == "7" {
			break
		} else {
			fmt.Println("Invalid Input")
		}
	}
}

func difficultyJava() {
	reader := bufio.NewReader(os.Stdin)
	var d string
	fmt.Println("||--------------------||")
	fmt.Println("Pick The Difficulty of Your Project!")
	fmt.Println("1. Beginner")
	fmt.Println("2. Intermediate")
	fmt.Println("3. Advanced")
	fmt.Println("4. 10x Developer (super-duper advanced)")
	fmt.Println("||--------------------||")
	fmt.Println("Input a number:")
	difficultyInput, _ := reader.ReadString('\n')
	d = strings.TrimSpace(difficultyInput)

	if d == "1" {
		javaBeginner()
	} else if d == "2" {
		javaIntermediate()
	} else if d == "3" {
		javaAdvanced()
	} else if d == "4" {
		javaDeveloper()
	} else {
		fmt.Println("Invalid Input")
	}
}

func difficultyGo() {
	reader := bufio.NewReader(os.Stdin)
	var d string
	fmt.Println("||--------------------||")
	fmt.Println("Pick The Difficulty of Your Project!")
	fmt.Println("1. Beginner")
	fmt.Println("2. Intermediate")
	fmt.Println("3. Advanced")
	fmt.Println("4. 10x Developer (super-duper advanced)")
	fmt.Println("||--------------------||")
	fmt.Println("Input a number:")
	difficultyInput, _ := reader.ReadString('\n')
	d = strings.TrimSpace(difficultyInput)

	if d == "1" {
		goBeginner()
	} else if d == "2" {
		goIntermediate()
	} else if d == "3" {
		goAdvanced()
	} else if d == "4" {
		goDeveloper()
	} else {
		fmt.Println("Invalid Input")
	}
}

func difficultyCPlus() {
	reader := bufio.NewReader(os.Stdin)
	var d string
	fmt.Println("||--------------------||")
	fmt.Println("Pick The Difficulty of Your Project!")
	fmt.Println("1. Beginner")
	fmt.Println("2. Intermediate")
	fmt.Println("3. Advanced")
	fmt.Println("4. 10x Developer (super-duper advanced)")
	fmt.Println("||--------------------||")
	fmt.Println("Input a number:")
	difficultyInput, _ := reader.ReadString('\n')
	d = strings.TrimSpace(difficultyInput)

	if d == "1" {
		cPlusBeginner()
	} else if d == "2" {
		cPlusIntermediate()
	} else if d == "3" {
		cPlusAdvanced()
	} else if d == "4" {
		cPlusDeveloper()
	} else {
		fmt.Println("Invalid Input")
	}
}

func difficultyJavascript() {
	reader := bufio.NewReader(os.Stdin)
	var d string
	fmt.Println("||--------------------||")
	fmt.Println("Pick The Difficulty of Your Project!")
	fmt.Println("1. Beginner")
	fmt.Println("2. Intermediate")
	fmt.Println("3. Advanced")
	fmt.Println("4. 10x Developer (super-duper advanced)")
	fmt.Println("||--------------------||")
	fmt.Println("Input a number:")
	difficultyInput, _ := reader.ReadString('\n')
	d = strings.TrimSpace(difficultyInput)

	if d == "1" {
		javascriptBeginner()
	} else if d == "2" {
		javascriptIntermediate()
	} else if d == "3" {
		javascriptAdvanced()
	} else if d == "4" {
		javascriptDeveloper()
	} else {
		fmt.Println("Invalid Input")
	}
}

func difficultyPython() {
	reader := bufio.NewReader(os.Stdin)
	var d string
	fmt.Println("||--------------------||")
	fmt.Println("Pick The Difficulty of Your Project!")
	fmt.Println("1. Beginner")
	fmt.Println("2. Intermediate")
	fmt.Println("3. Advanced")
	fmt.Println("4. 10x Developer (super-duper advanced)")
	fmt.Println("||--------------------||")
	fmt.Println("Input a number:")
	difficultyInput, _ := reader.ReadString('\n')
	d = strings.TrimSpace(difficultyInput)

	if d == "1" {
		pythonBeginner()
	} else if d == "2" {
		pythonIntermediate()
	} else if d == "3" {
		pythonAdvanced()
	} else if d == "4" {
		pythonDeveloper()
	} else {
		fmt.Println("Invalid Input")
	}
}

func difficultyCSharp() {
	reader := bufio.NewReader(os.Stdin)
	var d string
	fmt.Println("||--------------------||")
	fmt.Println("Pick The Difficulty of Your Project!")
	fmt.Println("1. Beginner")
	fmt.Println("2. Intermediate")
	fmt.Println("3. Advanced")
	fmt.Println("4. 10x Developer (super-duper advanced)")
	fmt.Println("||--------------------||")
	fmt.Println("Input a number:")
	difficultyInput, _ := reader.ReadString('\n')
	d = strings.TrimSpace(difficultyInput)

	if d == "1" {
		cSharpBeginner()
	} else if d == "2" {
		cSharpIntermediate()
	} else if d == "3" {
		cSharpAdvanced()
	} else if d == "4" {
		cSharpDeveloper()
	} else {
		fmt.Println("Invalid Input")
	}
}

func javaBeginner() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random beginner Java coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

func javaIntermediate() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random intermediate Java coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

func javaAdvanced() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random Advanced Java coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

func javaDeveloper() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random 10x developer Java coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

func javascriptBeginner() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random beginner Javascript coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

func javascriptIntermediate() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random Intermediate Javascript coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

func javascriptAdvanced() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random advanced Javascript coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

func javascriptDeveloper() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random 10x developer Javascript coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

func goBeginner() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random beginner Golang coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

func goIntermediate() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random Intermediate Golang coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

func goAdvanced() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random advanced Golang coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

func goDeveloper() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random 10x developer Golang coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

func cPlusBeginner() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random beginner C++ coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())
}

func cPlusIntermediate() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random intermediate C++ coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())
}

func cPlusAdvanced() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random C++ advanced coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

func cPlusDeveloper() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random 10x developer C++ coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

func cSharpBeginner() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random beginner C# coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

func cSharpIntermediate() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random intermediate C# coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

func cSharpAdvanced() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random advanced C# coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

func cSharpDeveloper() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random 10x developer C# coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

func pythonBeginner() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random beginner Python coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

func pythonIntermediate() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random intermediate Python coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

func pythonAdvanced() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random advanced Python coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

func pythonDeveloper() {
ctx := context.Background()
    client, err := genai.NewClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Give a random 10x developer Python coding project prompt with no code snippets in simple words with a quick explanantion"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())

}

