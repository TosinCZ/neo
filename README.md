# neo

# RaftKV

A distributed key-value database being built from scratch in Go, implementing
the Raft consensus algorithm, a custom storage engine, multi-node
replication, and automatic leader failover.

Built as a year-long project to go deep on distributed systems, concurrency,
and storage internals — see [`docs/PROCESS.md`](docs/PROCESS.md) for the
full log of decisions and difficulties along the way.

## Status

🚧 In progress (Subject to change!)
(started [October/2026]) 

- [ ] Single-node key-value store (in-memory + persistence)
- [ ] Custom storage engine (LSM-tree / B-tree)
- [ ] Multi-node replication (leader/follower, no fault tolerance yet)
- [ ] Raft consensus — leader election
- [ ] Raft consensus — log replication + safety
- [ ] Chaos testing / fault injection
- [ ] Web dashboard
- [ ] Benchmarks & writeup

## Why this project

Most student projects are CRUD apps. This one implements the actual hard
parts of a real distributed database — the same class of problem behind
systems like etcd and CockroachDB — including writing Raft myself rather
than using an existing consensus library.

## Architecture

```
neo/
├── cmd/           # entrypoints
├── internal/
├── docs/          # design docs, diagrams, process log
└── tests/          # integration / chaos tests
```


## Design decisions

Short version here; full reasoning and tradeoffs in
[`docs/PROCESS.md`](docs/PROCESS.md).

- **Language**: Go — chosen for strong concurrency primitives and faster iteration than Rust also I already have experience with it

## Benchmarks

*(To be added)*

## Demo

*(To be added)*

## License

MIT