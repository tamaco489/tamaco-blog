# Technology Stack

## Overview
This document describes the technology stack used in this project.

## Architecture Diagram

```mermaid
graph LR
    subgraph "Frontend"
        A[Next.js<br/>SSG + ISR]
    end

    subgraph "Hosting"
        B[AWS S3]
        C[CloudFront]
    end

    subgraph "Data Store"
        D[Supabase]
        E[DynamoDB<br/>Future]
    end

    subgraph "API Server"
        F[API Gateway]
        G[Lambda]
        H[ECR]
    end

    subgraph "CMS"
        I[Strapi<br/>Headless CMS]
    end

    subgraph "CI/CD"
        J[GitHub Actions]
    end

    subgraph "Monitoring"
        K[NewRelic]
        L[Google Analytics]
    end

    A --> B
    B --> C
    A --> F
    F --> G
    G --> H
    A --> I
    I --> D
    A --> D
    A --> L
    G --> K
```

## Details

### Hosting Server
- **AWS S3**: Static file hosting
- **AWS CloudFront**: CDN distribution
- **Next.js**: Static generation with SSG + ISR

### Data Store
- **Supabase**: Main database
- **DynamoDB**: Future implementation for scalability improvement

### API Server
- **API Gateway**: API endpoint management
- **Lambda**: Serverless functions
- **ECR**: Container image management

### Headless CMS
- **Strapi**: Content management system

### CI/CD
- **GitHub Actions**: Automated build and deployment

### Monitoring & Analytics
- **NewRelic**: Application monitoring
- **Google Analytics**: User behavior analytics

