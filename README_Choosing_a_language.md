# Day 1: Read up on different languages to get an overview
  Elm:
    @Paradigm - Functional
    @Typing - Static, Strong, Inferred
    @Designed-by - Evan Czaplicki
    @Resources:
      https://en.wikipedia.org/wiki/Elm_(programming_language)
      https://guide.elm-lang.org/
      https://www.youtube.com/watch?v=jrkLrm4Oh2s
      https://www.youtube.com/watch?v=7SacmHT7lLc

    @Pros:
      - Similar structure with every app.
      - "If you give Elm a shot and actually make a project in it, you will end up writing better JavaScript and React code."
      - The big reason to choose Elm over ClojureScript (CS) is types. Elm makes it really hard to write code that crashes, and it makes refactoring really easy, because the type checker helps you and shows where changes need to propagate to.
      - Elm is immutable, so there's no way to have state mutation accidentally hidden deep in your code. CS makes no such guarantees.
      - Elm has much better error messages and will aid you in finding bugs more easily (IMO) and much earlier in the development cycle. For this reason, it might be worth it to (at least) start with Elm.
    @Cons:
      Elm breaks backward compatibility almost every release.

  Elixer:
    @Paradigm - Functional, Concurrent, Distributed, Process-oriented
    @Typing - Dynamic, Strong
    @Resources:
      https://en.wikipedia.org/wiki/Elixir_(programming_language)
      http://elixir-lang.org/getting-started/introduction.html

    @Info:
      - Elixir runs on the BEAM, A.K.A the Erlang virtual machine and shares a lot in common with Erlang.

    @Other-Opinions:
      - Elixir looks like ruby on the surface, but is oh so different.
      - Because Elixir is a functional language, if you haven't dealt with one of those before, you'll be in for a bit of a learning curve.
      - Elixir is great for long-running stuff, like back-end web stuff.
      - It's concurrency (and ability to fully utilize multiple processors pretty much automatically) and fault tolerance stuff make it awesome for long running web apps.

  Go:
    @Paradigm - Functional, Concurrent, Logic, Imperative (object-based)
    @Typing - Static, Strong, Inferred, Structural
    @Developer - Google
    @Designed-by - Robert Griesemer, Rob Pike, Ken Thompson
    @Influenced-by - Alef, APL, BCPL, C, CSP, Limbo, Modula, Newsqueak, Oberon, occam, Pascal, Python, Smalltalk

    @Simplicity:
      *Pros*
        - Go is simple, relatively speaking. Its designers deliberately limited the scope of the language and built something that an experienced programmer can learn in an evening.
        - It’s easier to read code written by other Go programmers because it’s almost impossible for them to use a feature that you’ve forgotten or never learned.
      *Cons*
        - Go achieved its vaunted simplicity by leaving out dozens of clever, modern abstractions programmers are already using successfully.
        - Everyone has their favorite feature that didn’t make the Go cut. Some want generics; some want exceptions; some want more extensibility; some want something completely different.

    @Automation:
      *Pros*
        - Automatic declaration of variables.
      *Cons*
        - Automatic declaration can get details wrong when the same variable name is used in nested scopes.

    @Other-Opinions:
      *Pros*
        - It's a small language, pretty quick, and easy to learn.
        - Go is a fun imperative language and is great for small-to-mediumish programs (e.g., command line tool (CLI) projects).
        - It doesn't have the "startup cost" of a VM based language like Erlang/Elixir, so that's why I say it's great for CLI stuff, when we're comparing the two. (The BEAM (Erlang VM) definitely takes a noticeable amount of time to start up, so it's much more suited toward long running stuff.)
      *Cons*

    @Dane
      *Pros*
        He's a good source for questions.
      *Cons*
        I sometimes want to strangle him.

    @Resources:
      https://en.wikipedia.org/wiki/Go_(programming_language)
      http://www.infoworld.com/article/3146927/application-development/should-you-go-with-googles-go-7-pros-and-cons.html

  Clojure:
    @Paradigm - Functional
    @Typing - Dynamic, Strong
    @Designed-by - Rich Hickey
    @Influenced-by - C++, C#, Common Lisp, Erlang, Haskell, Mathematica, ML, Prolog, Scheme, Java, Racket, Ruby
    @Resources:
      https://en.wikipedia.org/wiki/Clojure

    @Pros:
      - Cljs has really nice tools and frameworks and very good performance.
      - CS is probably a bit more mature than Elm in terms of library support.
      - There are much more Clojure books (and other resources), so if along the way you feel that you really need a book to learn from, you might want to switch to Clojure.

    @Cons:
      - ClojureScript is dynamic and that made my days spent with it daunting to say the least. Most of the time I spent in ClojureScript was dedicated to endless debugging of the runtime errors like "'undefined' is not a function", which are freaking hard to trace down and appear every step of the way, once you try to abstract.
      - Overall the language makes it really hard to abstract, making your risk of hitting unpredictable runtime errors grow proportionally to the level of abstraction you reach out to.

  Other:
    https://smashingboxes.com/blog/choosing-your-future-tech-stack-clojure-vs-elixir-vs-go/
    https://www.reddit.com/r/elixir/comments/3c8yfz/how_does_go_compare_to_elixir/
    https://en.wikipedia.org/wiki/Functional_programming
    - Refresh Austin -- talked to a guy about his thoughts on Elm, Elixer, Go, and Clojure. Said Elm and Clojure were similar but preferred Elm. He also liked Go. Never used Elixer. Said to watch Rich Hickey videos on Elm.
    - Dane (Works in Go) - Asked his opinion on Go. Said there is a learning curve but it's not extreme. (When considering me and my experience) Suggests Elixer over Elm; suggests Go over Elixer. Dislikes Clojure outright. Suggests Elixer if I want to go fully down the functional programming route; even still, suggests Go over Elixer.

# Decision
  Go

  @Reasons --
    1) From my understanding, Go isn't really a Functional language. The others are functional languages. While I have nothing against functional languages, it may be a steeper learning curve than I want right now. And I don't want to get held up because I'm over ambitious.
    2) People claim Go is better for small-to-medium sized / command-line programs.
    3) People claim that an experienced programmer can learn in an evening
