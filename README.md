# Min Cost to Connect All Points

> NeetCode 150 · Advanced Graphs — MST on an implicit complete graph: dense Prim with a distance array beats the heap when every pair is an edge. Go, stdlib only, tested.

## Problem

Connect n plane points with wire priced at Manhattan distance so everything is reachable, minimum total cost — a minimum spanning tree where the edges are *implied* (all n² pairs exist).

## Approach

Heap-based Prim (dsa-168) would push O(n²) edges and pay a log factor for nothing. On a **complete** graph the O(n²) array form of Prim is both simpler and optimal: keep `dist[v]` = the cheapest wire from the current tree to each outside point; each round linearly scans for the cheapest outside point, absorbs it, and relaxes the remaining distances against the newcomer. No edge list is ever materialized — distances are computed on demand from coordinates. Choosing the array-Prim vs heap-Prim by graph density is the actual lesson; Kruskal would need all n² edges sorted and loses outright here.

**Complexity:** O(n²) time, O(n) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
