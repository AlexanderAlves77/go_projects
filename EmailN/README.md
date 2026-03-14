EmailN
│
├── cmd
│   └── api/
├── internal/
│   ├── contract/
│   │   ├── campaign_response.go
│   │   └── NewCampaignDto.go
│   │
│   ├── domain/
│   │    └── campaign/
│	│        ├── campaign_test.go
│	│        ├── campaign.go
│	│        ├── repository.go
│	│        ├── service_test.go
│   │        └── service.go
│   │
│   ├── infrastructure/
│   │   ├── database/
│	│   │	├── campaign_repository.go
│   │   │   └── new_db.go
│   │   │
│   │   └── mail/
│   │       └── send_email.go
│   │
│   ├── internal-errors/
│   │   ├── errors.go
│   │   └── validator.go
│   │
│   └── test/
│       └── internal-mock/
│           ├── campaign_repository_mock.go
│           └── campaign_service_mock.go
│   
├── .gitignore
├── go.mod
├── go.sum
└── README.md