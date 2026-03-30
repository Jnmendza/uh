# UnionHub

## Product Requirements Document

### The All-in-One Platform for Football Supporters Groups

| Version             | 1.0 — Initial Draft                          |
| :------------------ | :------------------------------------------- |
| **Date**            | [cite_start]March 2026 [cite: 4]             |
| **Status**          | [cite_start]In Review [cite: 4]              |
| **Author**          | [cite_start]UnionHub Founding Team [cite: 4] |
| **Confidentiality** | [cite_start]Internal Use Only [cite: 4]      |

---

## 1. Executive Summary

[cite_start]UnionHub is a B2B SaaS platform purpose-built for football (soccer) supporters groups[cite: 7]. [cite_start]It provides group administrators with a fully branded, all-in-one digital hub to manage their community—replacing fragmented tools like WhatsApp, Facebook Groups, and spreadsheets with a single, football-native platform[cite: 8]. [cite_start]UnionHub's initial launch will be a Progressive Web App (PWA), with native iOS and Android apps planned for a future phase[cite: 9].

## 2. Problem Statement

### 2.1 Current Landscape

[cite_start]Football supporters groups today rely on a patchwork of consumer tools not designed for their needs[cite: 15]:

- [cite_start]**WhatsApp & Telegram** groups for chat — no admin controls, no structure, chaotic at scale[cite: 16].
- [cite_start]**Facebook Groups** for announcements & events — declining engagement, algorithmic interference[cite: 17].
- [cite_start]**Spreadsheets or paper** for membership management — error-prone and time-consuming[cite: 18].
- [cite_start]**Third-party ticketing/event tools** — disjointed from the community[cite: 19].
- [cite_start]No centralized place for match stats, standings, chants, or merch[cite: 20].

### 2.2 Core Problems

- [cite_start]Supporters groups have no dedicated, football-specific platform built for their needs[cite: 22].
- [cite_start]Admins waste significant time managing members, events, and communications across multiple tools[cite: 23].
- [cite_start]Members experience a fragmented community with no single place to engage[cite: 24].
- [cite_start]There is no clean way to monetize membership tiers, sell merch, or manage giveaways[cite: 25].
- [cite_start]Football culture (chants, match days, stats) is completely absent from generic community tools[cite: 26].

## 3. Goals & Objectives

### 3.1 Business Goals

- [cite_start]Acquire the first paying supporters group customers within 3 months of launch[cite: 29].
- [cite_start]Validate the product with a real-world beta group (the founding team's own supporters group of 150+ members)[cite: 30].
- [cite_start]Build a scalable, multi-tenant SaaS platform[cite: 31].
- [cite_start]Establish UnionHub as the default platform for football supporters group management[cite: 32].

### 3.2 Product Goals

- [cite_start]Deliver an MVP covering high-priority features: Chat, Events, Membership, News, and Match Stats[cite: 34].
- [cite_start]Provide each group with a fully branded space (logo, colors, name)[cite: 35].
- [cite_start]Ensure the platform is usable by non-technical fans as a PWA[cite: 36].
- [cite_start]Integrate live football data via a sports API[cite: 38].

## 4. Target Users

- [cite_start]**Primary User: Supporters Group Admin / Organizer:** The paying customer responsible for managing the group and engaging the community[cite: 40, 41].
- [cite_start]**Secondary User: Group Member (Fan):** The end consumer who interacts with the community through the PWA[cite: 45, 46].
- [cite_start]**Platform User: Super Admin (UnionHub):** The platform owners managing all groups and subscriptions[cite: 50, 51].
- [cite_start]**Target Group Size:** Initially optimized for 50–500 members[cite: 53].

## 5. User Roles & Permissions

| Role            | Responsibilities                                                                                                  | Access Level                                          |
| :-------------- | :---------------------------------------------------------------------------------------------------------------- | :---------------------------------------------------- |
| **Super Admin** | [cite_start]Manages the UnionHub platform globally; creates group instances and oversees subscriptions[cite: 56]. | [cite_start]Full platform access [cite: 56]           |
| **Moderator**   | [cite_start]Assists group admin; manages content, chat, events, and news; cannot manage billing[cite: 56].        | [cite_start]Content & community management [cite: 56] |
| **Member**      | [cite_start]Standard member; views content, participates in chat, votes in polls, and RSVPs to events[cite: 56].  | [cite_start]Read + participate [cite: 56]             |

## 6. Platform & Technical Overview

### 6.1 Platform

- [cite_start]**Launch:** Progressive Web App (PWA)[cite: 59].
- [cite_start]**Future:** Native iOS & Android applications[cite: 60].
- [cite_start]**Architecture:** Multi-tenant architecture[cite: 61].

### 6.2 Branding & White-Labeling

[cite_start]Instances are customizable with group name, logo, brand colors, custom subdomains, and cover imagery[cite: 63, 64, 65, 66, 67].

### 6.4 Sports Data

[cite_start]Match stats and schedules powered by third-party APIs (e.g., API-Football, SportMonks)[cite: 74].

## 7. Feature Requirements

### PHASE 1 — MVP FEATURES

- [cite_start]**7.1 Chat Channels & Messaging (Priority #1):** Structured channels (e.g., #matchday), direct messaging, and moderator controls[cite: 87, 91, 92, 93].
- [cite_start]**7.2 Events (Priority #2):** Event creation, RSVP functionality, and calendar views linked to match fixtures[cite: 104, 106, 107, 110, 111].
- [cite_start]**7.3 Membership Packages (Priority #3):** Custom tiers (Free, Silver, Gold) with manual payment tracking at launch[cite: 119, 121, 125].
- [cite_start]**7.4 News & Updates (Priority #4):** Official group announcement feed with rich text and reactions[cite: 133, 135, 137].
- [cite_start]**7.5 Club Schedule & Match Stats (Priority #5):** Live scores, standings, and player statistics via API[cite: 141, 146, 147, 149].

### PHASE 2 — POST-MVP FEATURES

- [cite_start]**7.6 Polls:** Member voting on group decisions[cite: 154].
- [cite_start]**7.7 Gallery:** Member-submitted photos and videos[cite: 161, 163].
- [cite_start]**7.8 Giveaways:** Admin-run prizes and entries[cite: 168, 170].
- [cite_start]**7.9 Group Chants:** Digital songbook with lyrics and audio playback[cite: 175, 177, 178].
- [cite_start]**7.10 Member Directory:** Searchable member profiles[cite: 181, 183].
- [cite_start]**7.11 Merch Inventory & Sales:** Catalog browsing and manual order tracking[cite: 187, 189, 191].
- [cite_start]**7.12 Donations & Charity Campaigns:** Fundraising campaigns with live progress bars and raffle tie-ins[cite: 196, 200, 202].

## 8. Feature Summary

| Feature          | Description                        | Priority    | MVP?                        |
| :--------------- | :--------------------------------- | :---------- | :-------------------------- |
| Chat Channels    | Structured group chat & moderation | 🔴 Critical | [cite_start]Yes [cite: 219] |
| Events           | Creation, RSVP, and calendar       | 🔴 Critical | [cite_start]Yes [cite: 219] |
| Membership       | Tiered packages, manual tracking   | 🔴 Critical | [cite_start]Yes [cite: 219] |
| News & Updates   | Admin announcements feed           | 🟠 High     | [cite_start]Yes [cite: 219] |
| Club Schedule    | Live fixtures & stats via API      | 🟠 High     | [cite_start]Yes [cite: 219] |
| Polls            | Member voting                      | 🟡 Medium   | [cite_start]No [cite: 219]  |
| Gallery          | Photo/video submissions            | 🟡 Medium   | [cite_start]No [cite: 219]  |
| Giveaways        | Admin-run prizes                   | 🟡 Medium   | [cite_start]No [cite: 219]  |
| Group Chants     | Lyrics and audio songbook          | 🟡 Medium   | [cite_start]No [cite: 219]  |
| Member Directory | Searchable profiles                | 🟡 Medium   | [cite_start]No [cite: 219]  |
| Merch Sales      | Catalog and manual tracking        | 🟢 Low      | [cite_start]No [cite: 219]  |
| Donations        | Charity campaigns & progress bars  | 🟢 Low      | [cite_start]No [cite: 219]  |

## 9. Non-Functional Requirements

- [cite_start]**Performance:** PWA load time < 3 seconds; chat delivery < 1 second[cite: 222, 223].
- **Security:** Data encrypted at rest/transit; [cite_start]GDPR-compliant[cite: 226, 228].
- [cite_start]**Reliability:** 99.5% uptime target[cite: 232].
- [cite_start]**Scalability:** Support for up to 5,000 members per instance[cite: 240].

## 10. Phased Roadmap

- [cite_start]**Phase 1 — MVP (Months 1–4):** Core engagement, beta testing (150+ members), and branding[cite: 244, 245, 251].
- [cite_start]**Phase 2 — Growth (Months 5–8):** Polls, Gallery, Merch, and Native iOS/Android apps[cite: 254, 256, 257, 261, 262].
- [cite_start]**Phase 3 — Scale (Months 9–12):** Integrated Stripe payments, multi-language support, and advanced analytics[cite: 263, 265, 266, 267].

## 11. Success Metrics

- **MVP Success:** Beta group active within 60 days; [cite_start]70%+ weekly login rate[cite: 272, 273, 274].
- [cite_start]**Business KPIs:** MRR growth, MAU per group, and retention rate[cite: 277, 278, 279, 280].

## 12. Assumptions & Constraints

- [cite_start]**Assumptions:** Reliable sports API availability; admins are willing to pay for purpose-built tools[cite: 284, 285, 286].
- [cite_start]**Constraints:** Pricing model undefined; manual payments limit initial scalability[cite: 289, 290, 291].

---

**UnionHub — Confidential. For internal use only. [cite_start]© 2026 UnionHub.** [cite: 309]
