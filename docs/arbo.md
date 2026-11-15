
```
sma/                          (module "sma", go.mod)
├── agt/                      les agents + leurs stratégies
│   ├── agent.go              interface Agent (Start, Percept, Deliberate, Act)
│   ├── vehicle.go            VehicleAgent (1 goroutine par véhicule)
│   ├── coordinator.go        CoordinatorAgent (barrière, tri par ID, invariant)
│   ├── behavior.go           interface Behavior + RuleBehavior (règles ACN)
│   ├── behavior_llm.go       LLMBehavior + LLMBroker (extension, plan C)
│   ├── strategy.go           helpers de tri déterministe (départage par ID)
│   ├── strategy_equal.go     EqualShare
│   ├── strategy_edf.go       EarliestDeparture
│   ├── strategy_fcfs.go      FirstComeFirstServed
│   ├── strategy_llf.go       LeastLaxity
│   ├── strategy_sllf.go      SmoothedLLF
│   └── strategy_lowsoc.go    LowestSoCFirst
├── cmd/
│   ├── simu/                 une simulation (+ serveur web plus tard)
│   ├── bench/                N graines x stratégies -> CSV
│   ├── tst-agt/              test des agents
│   ├── tst-groq/             test du client Groq   (remplace tests/test_groq.go)
│   └── tst-jev/              test du client Jev    (remplace tests/test_jev.go)
├── env/                      environnement + types partagés
│   ├── types.go              VehicleID, DtHours, EVView, Allocation, PowerStrategy   <- les premiers types
│   ├── env.go                Environment (horloge, tick, WaitGroup)
│   ├── clock.go              temps simulé (1 tick = 5 min)
│   ├── site.go               Site, Charger, P_site(t)
│   ├── scenario.go           générateur à graine fixe + chargeur CSV ACN/Utrecht
│   ├── metrics.go            KPIs (calculés ici, jamais dans les agents)
│   ├── client.go             existant
│   ├── groq.go               existant
│   └── jev.go                existant
├── net/                      messages échangés entre agents
│   └── message.go            Tick, Request{View, Reply chan<- Grant}, Grant{KW}
├── web/                      front (plus tard) : index.html, app.js
└── go.mod
```
