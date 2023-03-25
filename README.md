# Kanban Board Application

This is a Kanban board application built with Golang, Fiber, Postgres, Vite, React and Tailwind CSS.

## What is a Kanban Board?

A Kanban board is a visual tool used to manage and track work in progress. It is based on the Kanban method, which was originally developed in Japan for use in manufacturing.

A typical Kanban board consists of a series of columns or sections that represent the stages of a process. Each column contains cards or sticky notes that represent work items, such as tasks or user stories. The cards are moved from one column to the next as the work progresses.

Kanban boards are useful for several reasons:

- They provide a clear, visual representation of the work that needs to be done and the progress that has been made.
- They help teams to prioritize work and focus on the most important tasks.
- They promote collaboration and communication by making it easy for team members to see what others are working on.
- They encourage continuous improvement by allowing teams to identify bottlenecks and areas for optimization.

## Requirements

- Docker
- Docker Compose

## Installation

1. Install Docker by following the instructions for your platform:

   - [Docker for Mac](https://docs.docker.com/docker-for-mac/install/)
   - [Docker for Windows](https://docs.docker.com/docker-for-windows/install/)
   - [Docker for Linux](https://docs.docker.com/engine/install/)

2. Install Docker Compose by following the instructions for your platform:

   - [Docker Compose for Mac](https://docs.docker.com/compose/install/)
   - [Docker Compose for Windows](https://docs.docker.com/compose/install/)
   - [Docker Compose for Linux](https://docs.docker.com/compose/install/)

3. Clone this repository

4. Develop and run the application using Docker Compose:

```
cd kanban-board
docker compose -f docker-compose.dev.yml up
```

This will start the dev environment in containers

## Desing Ideas

Once the application is running, you can use it to manage your Kanban board tasks.

1. Users can create boards for different projects or workflows, and invite others to collaborate on those boards.
2. Boards can be customized with columns, labels, due dates, and other features to fit the needs of each project or workflow.
3. Users can assign tasks to themselves or other team members, and set deadlines and priorities for those tasks.
4. Users can add comments, attachments, and checklists to tasks to provide additional context and information.
5. Boards can be organized into teams, which can have different levels of access and permissions. For example, team members might be able to view and edit all boards within the team, while guests might only be able to view certain boards.6.
6. Users can search for boards and tasks across their organization, and filter by various criteria such as due date, label, or team.
7. The application can provide analytics and reports to help users track progress and identify areas for improvement, such as bottlenecks or frequently overdue tasks.
8. Users can receive notifications and reminders about upcoming deadlines, new comments, or changes to tasks.
9. The application can integrate with other tools and services, such as email, calendar, or project management software, to provide a seamless workflow for users.
