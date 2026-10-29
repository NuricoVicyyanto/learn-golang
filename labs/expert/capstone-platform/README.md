# Capstone Lab — Production Platform

> **Level:** Expert · **Spec:** [Capstone Platform](../../../docs/03-expert/09-capstone-platform/README.md)

## Goal

Turn this multi-binary skeleton into the webhook delivery platform described in the capstone requirements.

## Layout

```text
capstone-platform/
├── cmd/api/             HTTP API and event ingestion
├── cmd/worker/          delivery workers
├── internal/config/     configuration
├── internal/domain/     tenants, events, deliveries
├── internal/store/      persistence boundary
├── internal/delivery/   signing, retry, HTTP delivery
└── deploy/k8s/          manifests to complete
```

## Run

```bash
cd labs
go run ./expert/capstone-platform/cmd/api
go run ./expert/capstone-platform/cmd/worker
go test ./expert/capstone-platform/... -v
```

The starter is in-memory so it compiles without infrastructure. Replace it with PostgreSQL, a queue, and real workers as the design requires.

## TODO order

1. Define tenant, endpoint, event, and delivery models.
2. Implement idempotent ingestion and durable storage.
3. Implement signed delivery with retries, jitter, and timeouts.
4. Add a dead-letter path and replay.
5. Add metrics, traces, structured logs, and SLOs.
6. Add Kubernetes manifests, HPA, PDB, and graceful shutdown.
7. Run a failure-injection test and a load test.
8. Write the architecture decision records and runbook.

## Definition of done

- [ ] Every functional requirement in the capstone spec is met.
- [ ] A failure injected between commit and publish loses no event.
- [ ] Retries are bounded and idempotent.
- [ ] DLQ and replay are demonstrated.
- [ ] Load test meets the stated target.
- [ ] Rollout under load drops no requests.
- [ ] Architecture review and postmortem are written.
