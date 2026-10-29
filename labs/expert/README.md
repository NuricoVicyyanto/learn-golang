# Expert Labs

Suggested folders:

```text
expert/
├── 01-runtime/
├── 02-performance/
├── 03-concurrency/
├── 04-distributed/
├── 05-grpc/
├── 06-security/
├── 07-data/
├── 08-kubernetes/
└── capstone-platform/
```

## Evidence over claims

Every expert exercise should produce an artifact you can show:

- A benchmark comparison (`benchstat` output).
- A CPU, heap, or goroutine profile.
- An execution trace with an explanation.
- A load test report with before/after numbers.
- A threat model with prioritized mitigations.
- An architecture decision record (ADR).
- A runbook and an incident postmortem.

## Reproducibility

- Pin tool versions in the README.
- Include the exact commands used.
- Record hardware or container limits for performance work.
- Keep generated profiles out of Git unless they are small and intentional.
- Use a local Kubernetes cluster (kind or minikube) and commit the manifests, not the cluster state.

## Capstone review

Before calling the platform complete:

- [ ] Failure injection demonstrates retries, DLQ, and replay.
- [ ] Load test meets the stated target.
- [ ] SLOs and alerts are documented.
- [ ] A rollout under load drops no requests.
- [ ] Security review has no unresolved critical finding.
- [ ] You can defend every trade-off in the architecture review.
