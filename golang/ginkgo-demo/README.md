```bash
PS C:\Users\hanwei\Documents\GitHub\example\golang\ginkgo-demo> go get github.com/onsi/ginkgo/ginkgo
go: downloading golang.org/x/tools v0.0.0-20201224043029-2b0845dc783e
go: downloading golang.org/x/sys v0.0.0-20210112080510-489259a85091
go: added github.com/fsnotify/fsnotify v1.4.9
go: added github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0
go: added github.com/onsi/ginkgo v1.16.5
go: added golang.org/x/sys v0.0.0-20210112080510-489259a85091
go: added golang.org/x/tools v0.0.0-20201224043029-2b0845dc783e
go: added gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7
PS C:\Users\hanwei\Documents\GitHub\example\golang\ginkgo-demo> go install github.com/onsi/ginkgo/ginkgo
PS C:\Users\hanwei\Documents\GitHub\example\golang\ginkgo-demo\books> ginkgo bootstrap
Generating ginkgo test suite bootstrap for books in:
        books_suite_test.go                         

Ginkgo 2.0 is coming soon!                                                                                                               
==========================                                                                                                               
Ginkgo 2.0 is under active development and will introduce several new features, improvements, and a small handful of breaking changes.   
A release candidate for 2.0 is now available and 2.0 should GA in Fall 2021.  Please give the RC a try and send us feedback!             
  - To learn more, view the migration guide at https://github.com/onsi/ginkgo/blob/ver2/docs/MIGRATING_TO_V2.md                          
  - For instructions on using the Release Candidate visit https://github.com/onsi/ginkgo/blob/ver2/docs/MIGRATING_TO_V2.md#using-the-beta
  - To comment, chime in at https://github.com/onsi/ginkgo/issues/711                                                                    
                                                                                                                                         
To silence this notice, set the environment variable: ACK_GINKGO_RC=true                                                                 
Alternatively you can: touch $HOME/.ack-ginkgo-rc                      
PS C:\Users\hanwei\Documents\GitHub\example\golang\ginkgo-demo\books> ginkgo 
Failed to compile books:

# ginkgo-demo/books
books_suite_test.go:7:2: missing go.sum entry for module providing package github.com/onsi/gomega (imported by ginkgo-demo/books); to add:
        go get -t ginkgo-demo/books
FAIL    ginkgo-demo/books [setup failed]
FAIL
Ginkgo ran 1 suite in 2.0274361s
Test Suite Failed
Ginkgo 2.0 is coming soon!
==========================
Ginkgo 2.0 is under active development and will introduce several new features, improvements, and a small handful of breaking changes.
A release candidate for 2.0 is now available and 2.0 should GA in Fall 2021.  Please give the RC a try and send us feedback!
  - To learn more, view the migration guide at https://github.com/onsi/ginkgo/blob/ver2/docs/MIGRATING_TO_V2.md
  - For instructions on using the Release Candidate visit https://github.com/onsi/ginkgo/blob/ver2/docs/MIGRATING_TO_V2.md#using-the-beta
  - To comment, chime in at https://github.com/onsi/ginkgo/issues/711

To silence this notice, set the environment variable: ACK_GINKGO_RC=true
Alternatively you can: touch $HOME/.ack-ginkgo-rc
PS C:\Users\hanwei\Documents\GitHub\example\golang\ginkgo-demo> go get -t ginkgo-demo/books
go: downloading golang.org/x/net v0.0.0-20201021035429-f5854403a974
PS C:\Users\hanwei\Documents\GitHub\example\golang\ginkgo-demo\books> ginkgo                     
Running Suite: Books Suite
==========================
Random Seed: 1653360859
Will run 0 of 0 specs


Ran 0 of 0 Specs in 0.003 seconds
SUCCESS! -- 0 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS

Ginkgo ran 1 suite in 4.079202s
Test Suite Passed
PS C:\Users\hanwei\Documents\GitHub\example\golang\ginkgo-demo\books> ginkgo generate books
  books_test.go


Ginkgo 2.0 is coming soon!
==========================
Ginkgo 2.0 is under active development and will introduce several new features, improvements, and a small handful of breaking changes.
A release candidate for 2.0 is now available and 2.0 should GA in Fall 2021.  Please give the RC a try and send us feedback!
  - To learn more, view the migration guide at https://github.com/onsi/ginkgo/blob/ver2/docs/MIGRATING_TO_V2.md
  - For instructions on using the Release Candidate visit https://github.com/onsi/ginkgo/blob/ver2/docs/MIGRATING_TO_V2.md#using-the-beta
  - To comment, chime in at https://github.com/onsi/ginkgo/issues/711

To silence this notice, set the environment variable: ACK_GINKGO_RC=true
PS C:\Users\hanwei\Documents\GitHub\example\golang\ginkgo-demo\books> cat books_test.go 
package books_test

  - To comment, chime in at https://github.com/onsi/ginkgo/issues/711

To silence this notice, set the environment variable: ACK_GINKGO_RC=true
Alternatively you can: touch $HOME/.ack-ginkgo-rc
PS C:\Users\hanwei\Documents\GitHub\example\golang\ginkgo-demo\books> cat .\books_test.go
package books_test

import (
        . "github.com/onsi/ginkgo"
        . "github.com/onsi/gomega"

        . "ginkgo-demo/books"
)

var _ = Describe("Books", func() {
        var (
                longBook  Book
                shortBook Book
        )

        BeforeEach(func() {
                longBook = Book{
                        Title:  "Les Miserables",
                        Author: "Victor Hugo",
                        Pages:  1488,
                }

                shortBook = Book{
                        Title:  "Fox In Socks",
                        Author: "Dr. Seuss",
                        Pages:  24,
                }
        })

        Describe("Categorizing book length", func() {
                Context("With more than 300 pages", func() {
                        It("should be a novel", func() {
                                Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
                        })
                })

                Context("With fewer than 300 pages", func() {
                        It("should be a short story", func() {
                                Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
                        })
                })
        })
})
PS C:\Users\hanwei\Documents\GitHub\example\golang\ginkgo-demo\books> ginkgo
Running Suite: Books Suite
==========================
Random Seed: 1653362237
Will run 2 of 2 specs

++
Ran 2 of 2 Specs in 0.005 seconds
SUCCESS! -- 2 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS

Ginkgo ran 1 suite in 1.3000423s
Test Suite Passed
PS C:\Users\hanwei\Documents\GitHub\example\golang\ginkgo-demo\books>

```