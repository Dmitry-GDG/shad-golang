### ПРОВЕРКИ:

  ```shell
  # Из папки программы:
  cd ..
  # Из корня репозитория:
  go run ./sum/main.go
  # Форматирование файлов по стандартам Go:
  go fmt ./sum/main.go
  # Проверка и импорт недостающих пакетов:
  goimports ./sum/main.go // goimports -w ./sum/main.go
  # Проверьте, что код проходит линтер:
  golangci-lint run ./sum/...
  # Проверьте, что ваше решение проходит тесты локально:
  go test ./sum/...
  ```
### TASKS

- group:    Hello World
  start:    11-02-2023 18:00
  deadline: 21-02-2023 23:59
  tasks:

    - task: ~~[sum](./sum/README.md)~~
      score: 100
    - task: ~~[tour0](./tour0/README.md)~~
      score: 100
    - task: ~~[wordcount](./wordcount/README.md)~~
      score: 100
    - task: ~~[urlfetch](./urlfetch/README.md)~~
      score: 100
    - task: ~~[fetchall](./fetchall/README.md)~~
      score: 100

- group:    Basics
  start:    18-02-2023 16:00
  deadline: 28-02-2023 23:59
  tasks:
    - task: ~~[hotelbusiness](./hotelbusiness/README.md)~~
      score: 100
    - task: [hogwarts](./hogwarts/README.md)
      score: 100
    - task: [utf8](./utf8/README.md)
      score: 100
    - task: [varfmt](./varfmt/README.md)
      score: 100
    - task: [speller](./speller/README.md)
      score: 100
    - task: [forth](./forth/README.md)
      score: 100

- group:    Interfaces
  start:    25-02-2023 16:30
  deadline: 07-03-2023 23:59
  tasks:
    - task: [otp](./otp/README.md)
      score: 100
    - task: [lrucache](./lrucache/README.md)
      score: 100
    - task: [externalsort](./externalsort/README.md)
      score: 100
    - task: [retryupdate](./retryupdate/README.md)
      score: 100
    - task: [ciletters](./ciletters/README.md)
      score: 100

- group:    Goroutines
  start:    04-03-2023 16:30
  deadline: 14-03-2023 23:59
  tasks:
    - task: [tour1](./tour1/README.md)
      score: 100
    - task: [once](./once/README.md)
      score: 100
    - task: [rwmutex](./rwmutex/README.md)
      score: 100
    - task: [waitgroup](./waitgroup/README.md)
      score: 100
    - task: [cond](./cond/README.md)
      score: 100
    - task: [ratelimit](./ratelimit/README.md)
      score: 100

- group:    "[HW] Gitfame"
  start:    04-03-2023 16:30
  deadline: 19-03-2023 23:59
  tasks:
   - task: [gitfame](./gitfame/README.md)
     score: 0

- group:    Testing
  start:    11-03-2023 13:00
  deadline: 21-03-2023 23:59
  tasks:
    - task: [testequal](./testequal/README.md)
      score: 100
    - task: [fileleak](./fileleak/README.md)
      score: 100
    - task: [tabletest](./tabletest/README.md)
      score: 100
    - task: [tparallel](./tparallel/README.md)
      score: 200
    - task: [iprange](./iprange/README.md)
      score: 100

- group:    Concurrency with shared memory
  start:    18-03-2023 15:59
  deadline: 29-03-2023 23:59
  tasks:
    - task: [dupcall](./dupcall/README.md)
      score: 200
    - task: [keylock](./keylock/README.md)
      score: 200
    - task: [batcher](./batcher/README.md)
      score: 200
    - task: [pubsub](./pubsub/README.md)
      score: 300

- group:    HTTP
  start:    25-03-2023 16:00
  deadline: 06-04-2023 23:59
  tasks:
    - task: [urlshortener](./urlshortener/README.md)
      score: 100
    - task: [digitalclock](./digitalclock/README.md)
      score: 100
    - task: [middleware](./middleware/README.md)
      score: 200
    - task: [olympics](./olympics/README.md)
      score: 200
    - task: [firewall](./firewall/README.md)
      score: 200

- group:    Generics
  start:    01-04-2023 16:00
  deadline: 11-04-2023 23:59
  tasks:
    - task: [genericsum](./genericsum/README.md)
      score: 100
    - task: [treeiter](./treeiter/README.md)
      score: 100
    - task: [coverme](./coverme/README.md)
      score: 300

- group:    SQL
  start:    08-04-2023 16:00
  deadline: 18-04-2023 23:59
  tasks:
    - task: [dao](./dao/README.md)
      score: 100
    - task: [ledger](./ledger/README.md)
      score: 200
    - task: [shopfront](./shopfront/README.md)
      score: 100
    - task: [wscat](./wscat/README.md)
      score: 200

- group:    Reflect
  start:    15-04-2023 16:00
  deadline: 25-04-2023 23:59
  tasks:
    - task: [reversemap](./reversemap/README.md)
      score: 100
    - task: [jsonlist](./jsonlist/README.md)
      score: 100
    - task: [jsonrpc](./jsonrpc/README.md)
      score: 100
    - task: [structtags](./structtags/README.md)
      score: 100

- group:    Low level
  start:    22-04-2023 16:00
  deadline: 03-05-2023 23:59
  tasks:
    - task: [illegal](./illegal/README.md)
      score: 100
    - task: [blowfish](./blowfish/README.md)
      score: 100

- group:    "[HW] Dist Build"
  start:    26-04-2023 12:00
  deadline: 29-05-2023 23:59
  tasks:
    - task: [distbuild](./distbuild/README.md)
      score: 0

- group:    Analysis
  start:    29-04-2023 18:00
  deadline: 10-05-2023 23:59
  tasks:
    - task: [testifycheck](./testifycheck/README.md)
      score: 200
    - task: [gzep](./gzep/README.md)
      score: 100
      
- group:    Bonus
  start:    11-02-2023 18:00
  deadline: 29-05-2023 23:59
  tasks:
    - task: [consistenthash](./consistenthash/README.md)
      score: 200
    - task: [gossip](./gossip/README.md)
      score: 300
    - task: [smartsched](./smartsched/README.md)
      score: 200
      watch:
        - distbuild/pkg/scheduler
    - task: [wasm](./wasm/README.md)
      score: 300
