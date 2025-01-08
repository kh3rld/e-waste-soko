# E-Waste Soko

E-Waste Soko is a community-driven platform designed to address the pressing issue of electronic waste. By facilitating convenient e-waste collection and promoting responsible recycling practices, E-Waste Soko empowers individuals and organizations to contribute to environmental conservation while earning rewards.

## Features

- **Scheduled E-Waste Collection**: Users can schedule pickups through the platform.
- **Rewards System**: Participants earn tokens redeemable for essential goods or services.
- **Educational Resources**: Promotes awareness of e-waste disposal practices.
- **Impact Reporting**: Offers transparency on the environmental benefits of recycling efforts.

## Technology Stack

- **Backend**: Go
- **Frontend**: HTML, CSS, JavaScript
- **Database**: PostgreSQL
- **Containerization**: Docker
- **Authentication**: Firebase
- **Messaging**: Twilio API

## Challenges

- Designing a user-friendly interface.
- Ensuring reliable scheduling and logistics for e-waste collection.
- Integrating various APIs for authentication and messaging.
- Managing database performance with real-time data updates.

## Accomplishments

- Developed a seamless e-waste scheduling system.
- Successfully implemented a rewards mechanism.
- Established partnerships with certified recycling organizations.

## Roadmap

1. Enhance the rewards system with more partner vendors.
2. Introduce gamification elements to increase user engagement.
3. Expand educational resources to include video tutorials and infographics.
4. Integrate AI to predict peak collection times for optimized logistics.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/kh3rld/e-waste-soko.git
   cd e-waste-soko
   ```
2. Set up Docker and PostgreSQL:

### Docker Setup

1. Ensure Docker is installed on your machine. [Download Docker](https://www.docker.com/products/docker-desktop)
2. Start Docker and build the application containers:
   ```bash
   docker-compose up --build
   ```
3. This will create containers for the backend, frontend, and PostgreSQL database.

### PostgreSQL Setup

1. Access the PostgreSQL container:
   ```bash
   docker exec -it postgres-container-name psql -U postgres
   ```
2. Create the database:
   ```sql
   CREATE DATABASE soko_portal;
   ```
3. Run and apply database migrations:
   ```bash
   go run backend/cmd/main.go
   ```

## Usage

1. Navigate to the frontend URL (http://localhost:3000) to access the application.
2. Sign up, schedule pickups, and track your rewards.

## Contributing

Contributions are welcome! 

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

