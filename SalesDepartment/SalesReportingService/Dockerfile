# Use Node.js official image as a base
FROM node:18-alpine

# Set working directory in container
WORKDIR /usr/src/app

# Copy package.json and install dependencies
COPY package.json .

RUN npm install

# Copy the source code
COPY . .

# Expose the port the app runs on
EXPOSE 3000

# Command to run the app
CMD ["node", "index.js"]
