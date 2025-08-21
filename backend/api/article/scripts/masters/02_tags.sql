-- Tags master data
INSERT INTO "tags" ("name", "slug") VALUES
-- Programming Languages
('Go', 'golang'),
('TypeScript', 'typescript'),
('JavaScript', 'javascript'),
('Python', 'python'),
('Rust', 'rust'),
('Java', 'java'),
('PHP', 'php'),
('SQL', 'sql'),

-- Frameworks & Libraries
('React', 'react'),
('Next.js', 'nextjs'),
('Vue.js', 'vuejs'),
('Node.js', 'nodejs'),
('Express', 'express'),
('FastAPI', 'fastapi'),
('Django', 'django'),
('Spring Boot', 'spring-boot'),
('Gin', 'gin'),

-- Frontend Technologies
('HTML', 'html'),
('CSS', 'css'),
('Sass', 'sass'),
('TailwindCSS', 'tailwindcss'),
('Webpack', 'webpack'),
('Vite', 'vite'),

-- Databases
('PostgreSQL', 'postgresql'),
('MySQL', 'mysql'),
('MongoDB', 'mongodb'),
('Redis', 'redis'),
('SQLite', 'sqlite'),

-- DevOps & Infrastructure
('Docker', 'docker'),
('Kubernetes', 'kubernetes'),
('AWS', 'aws'),
('GCP', 'gcp'),
('Azure', 'azure'),
('Terraform', 'terraform'),
('GitHub Actions', 'github-actions'),
('CI/CD', 'cicd'),

-- Development Tools
('Git', 'git'),
('VS Code', 'vscode'),
('Linux', 'linux'),
('API', 'api'),
('REST API', 'rest-api'),
('GraphQL', 'graphql'),
('gRPC', 'grpc'),

-- Architecture & Design
('Clean Architecture', 'clean-architecture'),
('DDD', 'domain-driven-design'),
('Microservices', 'microservices'),
('Event Sourcing', 'event-sourcing'),
('CQRS', 'cqrs'),

-- Testing
('Unit Testing', 'unit-testing'),
('Integration Testing', 'integration-testing'),
('E2E Testing', 'e2e-testing'),
('Jest', 'jest'),
('Cypress', 'cypress'),

-- Concepts
('Performance', 'performance'),
('Security', 'security'),
('Monitoring', 'monitoring'),
('Logging', 'logging'),
('Caching', 'caching'),
('Optimization', 'optimization'),
('Debugging', 'debugging'),
('Refactoring', 'refactoring')

ON CONFLICT ("slug") DO UPDATE SET
    "name" = EXCLUDED."name",
    "updated_at" = CURRENT_TIMESTAMP;
