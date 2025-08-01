Here's a simple, step-by-step guide to get a new Go project up and running. All commands assume you're on a Unix-style shell (macOS, Linux, Git Bash on Windows), but they work on Windows PowerShell too with minor tweaks.

1. Install Go

   - Download the latest Go installer from [https://go.dev/dl](https://go.dev/dl)
   - Follow the installer instructions for your OS
   - Verify installation by running:
     ```bash
     go version
     ```
     You should see something like:
     ```
     go version go1.21.0 darwin/amd64
     ```

2. Create a project directory

   ```bash
   mkdir hello-world
   cd hello-world
   ```

3. Initialize a Go module Go modules manage dependencies and versioning. Choose a module path that matches where you’ll host code (for now you can use `nihini/hello-world` or your own repo URL):

   ```bash
   go mod init nihini/hello-world
   ```

   This creates a `go.mod` file with your module name.

4. Add a main file Create a file named `main.go` with a simple "Hello, world" program:

   ```go
   package main

   import "fmt"

   func main() {
       fmt.Println("Hello, world")
   }
   ```

5. Run your program

   ```bash
   go run .
   ```

   You should see:

   ```
   Hello, world
   ```

6. Build an executable To compile a binary you can run anywhere:

   ```bash
   go build -o hello-world
   ```

   This produces an executable named `hello-world` (or `hello-world.exe` on Windows).

7. Add dependencies When you import a third-party package, for example:

   ```go
   import "github.com/google/uuid" 
   ```

   Then run:

   ```bash
   go get github.com/google/uuid
   ```

   Your `go.mod` and `go.sum` will be updated automatically.

8. Organize your code As your project grows, use subdirectories:

   ```
   hello-world/
     cmd/           # entry points (sub-commands or apps)
       hello-world/   # contains main.go
     pkg/           # shared library code
     internal/      # private code not exposed as module API
     go.mod
     go.sum
   ```

9. Test your code Write tests alongside your code in `_test.go` files. For example, in `pkg/foo/foo.go`:

   ```go
   package foo

   func Add(a, b int) int {
       return a + b
   }
   ```

   And test in `pkg/foo/foo_test.go`:

   ```go
   package foo

   import "testing"

   func TestAdd(t *testing.T) {
       if Add(2, 3) != 5 {
           t.Fatal("expected 5")
       }
   }
   ```

   Run tests with:

   ```bash
   go test ./...
   ```

10. Version control Initialize git and make your first commit:

    ```bash
    git init
    # Create a public or private repo on Github and then connect it here: 
    git remote add origin git@github.com:nihini/hello-world.git
    git add .
    git commit -m "Initial commit"
    git push -u origin main
    ```

That’s it. You now have a basic Go project, with module support, build and test commands, and a structure ready to grow.

