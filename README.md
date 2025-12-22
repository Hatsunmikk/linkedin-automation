# linkedin-automation

LinkedIn Automation – Go (Proof of Concept)

Important Disclaimer – Educational Use Only.
This project is a technical proof-of-concept created solely for evaluation and learning purposes.
Automating LinkedIn violates LinkedIn’s Terms of Service.
This tool must never be used on real LinkedIn accounts or deployed in production.

## Demo Video

Watch the full walkthrough here:  
https://youtu.be/mCVrr0nB8hA

# Project Overview

This repository contains a Go-based browser automation proof-of-concept built using the Rod library.
The project demonstrates automation concepts, human-like behavior simulation, and anti-bot detection techniques while maintaining clean architecture and safety constraints.

The primary goal is to showcase:

- Browser automation design patterns

- Anti-detection and stealth techniques

- Modular Go code organization

- State persistence and resumable workflows

All real LinkedIn interactions are intentionally simulated.

# Key Features

1. Authentication System

- Credentials loaded securely from environment variables

- Graceful handling of missing credentials

- Detection of security checkpoints (e.g. captcha / 2FA)

- Fail-fast behavior to prevent unsafe execution

2. Search & Targeting (PoC-safe)

- Structured search queries (job title, company, location, keywords)

- Pagination handling across result pages

- DOM-based profile URL extraction

- Automatic duplicate detection

- Mock DOM injection for deterministic results (no real LinkedIn automation)

3. Connection Requests (PoC-safe)

- Simulated navigation to profile pages

- Personalized connection notes via templates

- Daily request limits enforced

- Duplicate request prevention

- All actions recorded in persistent state

4. Messaging System (PoC-safe)

- Detection of accepted connections (simulated)

- Automated follow-up messaging

- Template-based message generation

- Message history tracking

- Duplicate message prevention

# Anti-Bot Detection & Stealth Techniques

This project implements 8 human-like stealth mechanisms, including all mandatory requirements:

Mandatory :

- Human-like mouse movement (curved paths, overshoot, micro-corrections)

- Randomized timing patterns (think time, delays, pacing)

- Browser fingerprint masking (viewport randomization, disabling automation flags)

Additional Techniques :

- Random scrolling behavior

- Realistic typing simulation (variable speed, corrections)

- Mouse hovering and cursor wandering

- Activity scheduling (business hours only)

- Rate limiting and throttling of actions

These techniques are implemented to demonstrate understanding of bot detection vectors, not to bypass real-world protections.

# Architecture Overview

The project follows a modular, maintainable Go architecture:

cmd/app/
main.go # Application entry point

internal/
auth/ # Authentication logic
browser/ # Browser lifecycle management
config/ # Configuration loading & validation
connections/ # Connection request workflow
messaging/ # Follow-up messaging system
search/ # Search & targeting engine
stealth/ # Anti-detection mechanisms
state/ # Persistent state management
logger/ # Structured logging

Each package has a single responsibility and well-defined boundaries.

# State Persistence

The application maintains a persistent JSON-based state file (state.json) to track:

- Sent connection requests

- Accepted connections (simulated)

- Sent follow-up messages

This allows:

- Safe resumption after interruption

- Duplicate prevention

- Deterministic behavior across runs

# Configuration & Environment Variables

.env.example ;

LINKEDIN_EMAIL=your_email_here
LINKEDIN_PASSWORD=your_password_here

Environment variables are loaded using godotenv for local development.

Never commit real credentials.
.env is ignored via .gitignore.

# Running the Project

Prerequisites :

- Go 1.20+

- Windows / macOS / Linux

- Internet connection (for Chromium download via Rod)

Run:
go run cmd/app/main.go

The automation flow will:

1. Load configuration

2. Launch Chromium

3. Authenticate (env-based)

4. Execute stealth behaviors

5. Run search -> connections -> messaging (PoC-safe)

6. Persist state and exit cleanly

# Demonstration Video

A walkthrough video is included / linked in this repository demonstrating:

- Project setup

- Configuration

- Execution flow

- Key automation and stealth features

(See video link in repository.)

# Design Decisions & Trade-offs

- No real LinkedIn automation: All interactions are simulated to respect ToS

- Mock DOM injection: Ensures deterministic behavior and safe testing

- State-first design: Prevents duplicates and enables resumability

- Separation of concerns: Browser logic is isolated from business logic

- Version-safe Rod usage: Compatible with Rod v0.116+

# Evaluation Alignment

This project directly addresses the assignment evaluation criteria:

- Anti-Detection Quality: Comprehensive stealth mechanisms

- Automation Correctness: End-to-end simulated workflow

- Code Architecture: Modular, idiomatic Go design

- Practical Implementation: Realistic patterns with safe constraints

# Author

Debosmita Banerjee
