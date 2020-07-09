<!-- Attribution-NonCommercial-NoDerivs 2.5
   - https://spdx.org/licenses/CC-BY-NC-ND-2.5.html
   - (C) 2020 FreightTrust and Clearing Corporation -->

# Freight Trust & Clearing Omnibus

!!! abstract The Omnibus format of documentation includes a **Rulebook**,
**Engineering Specifications**, **Technical Documentation**, **Research
Papers**, and more. [Freight Trust & Clearing](https://freighttrust.com).

- [Freight Trust & Clearing Omnibus](#freight-trust---clearing-omnibus)
  - [Overview](#overview)
  - [Rulebook Notice](#rulebook-notice)
  - [Overview](#overview-1)
  - [Freight Trust Network Overview](#freight-trust-network-overview)
    - [Protocol](#protocol)
  - [Nodes and Nodlets](#nodes-and-nodlets)
    - [Nodes](#nodes)
    - [Nodlet](#nodlet)
      - [Sub Node Classes](#sub-node-classes)
  - [Threshold Utility Level](#threshold-utility-level)
  - [Categories for Network Operations](#categories-for-network-operations)
  - [End-market fundamentals](#end-market-fundamentals)
    - [Governance](#governance)

!!! note
[For a system design model, vist our design page here](https://freight-trust-model.netlify.app/)

### Resources

Quick System Scope

> This is what it takes for a transaction to occur

- UserRoles
- UserTypes
- Container
- ContainerCodes (STIFL)
- Conveyance
  - Non-negotiable
    - Bill of Lading
    - Warehouse Receipt
  - Negotiable
    - Bill of Lading
    - Warehouse Receipt
- Transactions
  - Electronic Data Interchange
    - X12
      - 4010 204 Motor Carrier Shipment Information 210 Freight Details and
        Invoice 211 Motor Carrier Bill of Lading 214 Shipment Status Message 997
        Functional Acknowledgement
        - Additional Transactions Supported
- Protocol
  - AS2 Protocol
  - Concenus Protocol
  - MFT Protocol
  * PKI
    - x.509
    - ECDSA
  - Maidenlane

* [Network Design](/concepts/network-design.md)
* [Ecosystem](/concepts/ecosystem.md)
* [Protocol](/concepts/as2protocol.md)
* [Pricing Precedence](/operations/pricing-spec-precedence.md)
* [Pricing Benchmark](/operations/pricing-benchmark.md)

## Overview

```
Guide to the documentation is as follows:
    * Category *an informal name given to a section or collection of sections*.
       - Section *a formal name given to a specific part of the* `omnibus`.
           + Sub-Section *a formal name given to a specific part of a* `section`.
```

- Operations

* Corporate ...- Pricing Benchmark - Pricing Precedence - Transaction Pricing -
  Incident Response Plan - Software Defects - Business Continuity

- Explainations

* Concepts
  - [Mechanics & Dogma](concepts/central-dogma.html): This can be thought of as
    an abbreviated "informal academic paper"

- Legal

* Rulebook
  - [General](#): Information on the legal paramters that constitute how
    behavior is regulated on-chain.

- Technical

* Blockchain

  - Token Overview: Information on the \$EDI Token and its utility
  - Staking: Information on how Staking Pools
  - Network: Information on "master nodes" and the "pools" for nodes

* EDI (X12)
  - Transactions with available Parsing and Mapping Support
  - Available Transactions with just Parsing
* EDI General
  - All supported Transactions
  - Schema Registry

## Rulebook Notice

!!! note In these Rules, unless the context clearly requires otherwise,

    (a) words in the singular include the plural and words in the plural include the singular,
    (b) any gender includes each other gender,
    (c) references to statutory provisions include those provisions, and any rules or regulations promulgated thereunder, as amended, and.
    (d) all uses of the word “including” should be construed to mean “including, but not limited to.


    Headings included herein are for convenience purposes only and do not form a part of these Rules.

!!! important **Date and Time References** Unless otherwise specified, all
references. to dates, times or time periods shall refer to, or be measured in.
accordance with the time in New York City, New York.

??? note "Purpose" The modern digital business works in real time; it informs
interested parties of things of interest when they happen, it makes sense of,
and derives insight from an ever-growing number of sources. It learns, predicts
and is intelligent -- it is by nature Event Driven.Event-driven architecture
(EDA) is an architecture pattern that promotes the production, detection,
consumption of, and reaction to events. This architectural pattern can be
applied to the systems that transmit events among loosely coupled software
components and services.

## Overview

- Lower transaction cost to drive evolution in markets
- Look for fragmented markets as an opportunity to seek for efficiencies
- Pass on to the customer the savings introduced by efficiencies
- Innovate pricing thanks to marketplace mechanism and data
- Build mechanisms that support suppliers of all size
- Invest in creating a mechanism that brings the best to the top (i.e. rate of
  failure)
- Invest in mechanisms for empowerment, augmenting workers potential
  [Read more ...](corporate/README.md)

!!! info Event Managed State Real time insights/intelligence

- Being able to communicate and persist events.
- Being able to take direct action on events.
- Processing event streams to derive real time insight/intelligence.
- Providing communication for event driven microservices.

## Freight Trust Network Overview

We adopt a “cryptoasset” pricing model, evaluate our current token model, and
propose several changes in order to maximize value accrual as the network opens
for public (i.e. business) usage. Much of this work is based on sources found at
the end, and we would like to thank everyone who offered their critiques and
thoughts on the subject.

Examining the “bifurcated” token model (having 2 tokens: a transactional token
(network native) and a “system” token (erc-20) is shown to be a better model
than a traditional singular token model.

#### Protocol

> On Protocols

    > "Protocols encode the rules of engagement that coordinate the
    > exchange of a service between a global supplier and global
    > consumer."
    >
    > —  Chris Burniske Protocols As Minimally Extractive Coordinators

Protocols provide structure for businesses, but are not businesses themselves;
they are systems of logic that coordinate exchange between suppliers
(businesses) and consumers of a service. As coordinators of exchange, protocols
should be minimally extractive, whereas businesses are incentivized to be
maximally extractive (that’s profit, and a business is valued as a multiple of
its profit). The ERC-20 token, \$EDI, acts as a right to receive the native
network asset, xEDI on demand. xEDI does not act as a right to receive any other
asset, instead, it grants the right to transact on the network-based upon market
conditions and pricing. This does not preclude a trading mechanism from
existing: we simply enforce a one-way mechanism for EDI to xEDI.

## Nodes and Nodlets

<div class="tip">

1 Validator has 3 Nodlets making the total of 4 "containers"/"servers"

</div>

| Gmemory              | 3        | Paid for every additional word when expanding memory.                      |
| -------------------- | -------- | -------------------------------------------------------------------------- |
| Gtxcreate            | 32000    | Paid by all contract-creating transactions after the Homestead transition. |
| Gtxdatazero          | 4        | Paid for every zero byte of data or code for a transaction.                |
| Gtxdatanonzero       | 68       | Paid for every non-zero byte of data or code for a transaction.            |
| Gtransaction         | 21000    | Paid for every transaction.                                                |
| ext4                 | 4096     | bytes                                                                      |
| EDI                  | 1        | message                                                                    |
| Tthroughput          | 135      | tx/s                                                                       |
| Block Size           | 8294400  |                                                                            |
|                      | 552960   | bytes per block                                                            |
|                      | 552.96   | kilo bytes per block                                                       |
| checkpoint increment | 256      |                                                                            |
| Rocksdb              | 3860     |                                                                            |
| epoch                | 22118400 |                                                                            |

### Nodes

Masternodes == Nodes. All nodes are technically the same. The classifications
below are more specifically `slots`. Nodes are rotated into different
`slot positions` over an `epoch`. An `epoch` is defined in terms of
`checkpoint increment` authority: concensus signing node validator: failover

### Nodlet

There are 4 maine classes of `nodes`. A `master-node` and the 3 `nodelets`

#### Sub Node Classes

- besu-tx: Handling of local transaction pool.
- besu-sync: Handling of blockchain synchronisation through Ethereum P2P
  network.
- besu-query: Handling of database queries.

## Threshold Utility Level

The `ERC-20` token, `$EDI`, can enable access to the network level operator
system, commonly referred to as "master nodes". This is both a network and a
legal operation as operators are part of a limited partnership that is
compensated from Freight Trust Inc.

Network operation is distributed amongst “Master Nodes”. An amount of \$EDI is
required, x, in order to be able to be eligible for hosting part of the main
network.

## Categories for Network Operations

```json
[
  { "Category": "Nodes", "Item": "Authority,Validator,Failover" },
  { "Category": "Nodlet", "Item": "query,sync,transaction" },
  { "Category": "Staking", "Item": "Operator,Pool" }
]
```

## End-market fundamentals

In our use case, we assume that the platform will accrue transaction volumes at
a consistent dollar rate per user transaction set (e.g. EDI transaction to a
trading partner). This transaction volume would likely be composed of an
end-user directly paying for the transaction on a per kilobyte basis.

**Value flows being distributed in the native cryptoasset requires an exogenous
assumption to form a base of value**, and even then, the supply-side’s
assessment of the value they’re receiving can be unpredictable given the native
asset’s volatility. We can make the assumption that at equilibrium the service
should be 1/10th the cost of a service provisioned by a company comparable
(e.g., storage at 1/10th the cost of AWS), as only networks that are more
efficient than companies will be able to get demand-sides to make the leap. The
cost of the service can then be pinpointed, the units of the service consumed
can be projected, and the product of the two is the value stream flowing to the
supply-side.

That value stream is what the cryptocapital has claim to, and so the capital
asset component of value can be assessed. The “outside” anchor is the cost of
comparable services that the cryptonetwork (freight trust network) will be
competing with, though trouble remains if the native cryptoasset is too
volatile, as that volatility influences the supply-sides’ perception of what
they’re getting paid. Freight Trust, Inc includes a few “legacy” mechanisms to
address this point, such as payment to node operators in fiat, thus establishing
a fixed cost floor.

### Governance

Governance mechanisms in general should: 1\) Attract a willing supply-side that
produces the resource (node operators, stakeholders, community et al)

2\) Connect that resource to a demand-side that values and is willing to pay for
the resource (businesses)

3\) It creates an open layer of access to the underlying resource for
distributors that want to build the last mile to the consumer. (Freight Trust
Corporate) As of right now, no formal governance mechanism besides the informal
RFC process exists for EDI. We plan in the following days to formally propose a
new governance mechanism based upon Dynamic Alignment. We have included a
graphic to illustrate the idea behind it
