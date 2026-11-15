# Critical Path Method (CPM) & Topological Wave DAG Task Scheduler (`topological_task_dag`)

> Category: **Graph Scheduling** | Language: **Go 1.22+**

## Overview

Schedules weighted dependency DAGs via Kahn wave layering, computing Earliest Start, Latest Start, total makespan, zero-slack critical tasks, and parallel execution waves.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(V + E) Kahn waves, forward ES/EF, backward LS/LF pass` |
| **Space Complexity** | `O(V + E) forward & reverse adjacency` |

## Mathematical & Safety Invariants

- Every edge u -> v satisfies EarlyStart[v] >= EarlyFinish[u] == EarlyStart[u] + Duration[u].
- On valid DAGs, all critical path nodes have zero float (LatestStart == EarlyStart) and positive makespan.

## Module Architecture

| File | Lifecycle Role |
| :--- | :--- |
| `types.go` | Stage 01 |
| `errors.go` | Stage 02 |
| `interface.go` | Stage 03 |
| `core.go` | Stage 04 |
| `validation.go` | Stage 05 |
| `fast_path.go` | Stage 06 |
| `concurrent.go` | Stage 07 |
| `invariants.go` | Stage 08 |
| `unit_test.go` | Stage 09 |
| `property_test.go` | Stage 10 |
| `benchmark_test.go` | Stage 11 |
| `README.md` | Stage 12 |

## Verification & Execution Example

```text
go test -v ./src/modules/topological_task_dag
```

---
© 2026 Igor Kan
