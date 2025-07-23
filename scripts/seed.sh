#!/bin/bash

echo "Seeding database with 100 facts..."

# Run the seeding script
go run scripts/seed_facts.go

echo "Database seeding completed!"
