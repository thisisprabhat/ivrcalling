# MongoDB Setup Guide for IVR Calling System

This guide covers MongoDB setup options for the IVR Calling System.

## Option 1: Local MongoDB (Development)

### macOS

1. **Install MongoDB using Homebrew:**
```bash
brew tap mongodb/brew
brew install mongodb-community
```

2. **Start MongoDB:**
```bash
# Start as a service
brew services start mongodb-community

# Or run directly
mongod --config /usr/local/etc/mongod.conf
```

3. **Verify MongoDB is running:**
```bash
mongosh
# You should see MongoDB shell
```

4. **Update .env:**
```env
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=ivr_calling_system
```

### Linux (Ubuntu/Debian)

1. **Install MongoDB:**
```bash
wget -qO - https://www.mongodb.org/static/pgp/server-7.0.asc | sudo apt-key add -
echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu $(lsb_release -cs)/mongodb-org/7.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-7.0.list
sudo apt-get update
sudo apt-get install -y mongodb-org
```

2. **Start MongoDB:**
```bash
sudo systemctl start mongod
sudo systemctl enable mongod
```

3. **Verify:**
```bash
sudo systemctl status mongod
mongosh
```

### Windows

1. **Download MongoDB:**
   - Visit [mongodb.com/try/download/community](https://www.mongodb.com/try/download/community)
   - Download Windows installer
   - Run installer (install as a service)

2. **MongoDB will run automatically**

3. **Verify:**
```cmd
mongosh
```

### Docker (All Platforms)

1. **Run MongoDB in Docker:**
```bash
docker run -d \
  --name mongodb \
  -p 27017:27017 \
  -v mongodb_data:/data/db \
  mongo:7.0
```

2. **Verify:**
```bash
docker ps
mongosh
```

## Option 2: MongoDB Atlas (Cloud - Recommended for Production)

### Setup

1. **Create Account:**
   - Go to [mongodb.com/cloud/atlas](https://www.mongodb.com/cloud/atlas)
   - Sign up for free account
   - Create a new project

2. **Create Cluster:**
   - Click "Build a Database"
   - Choose **FREE** tier (M0)
   - Select region closest to you
   - Click "Create Cluster"

3. **Configure Database Access:**
   - Go to "Database Access"
   - Click "Add New Database User"
   - Create username and password
   - Grant "Read and Write" privileges
   - Click "Add User"

4. **Configure Network Access:**
   - Go to "Network Access"
   - Click "Add IP Address"
   - For development: Click "Allow Access from Anywhere" (0.0.0.0/0)
   - For production: Add your server IP
   - Click "Confirm"

5. **Get Connection String:**
   - Go to "Database" → "Connect"
   - Choose "Connect your application"
   - Copy the connection string
   - Replace `<password>` with your database user password

6. **Update .env:**
```env
MONGODB_URI=mongodb+srv://username:password@cluster0.xxxxx.mongodb.net/?retryWrites=true&w=majority
MONGODB_DATABASE=ivr_calling_system
```

### MongoDB Atlas Advantages

✅ **Free tier available** (512 MB storage)  
✅ **Automatic backups**  
✅ **High availability**  
✅ **Global deployment**  
✅ **No server management**  
✅ **Built-in monitoring**

## Verifying Connection

### Using mongosh

```bash
# Local
mongosh mongodb://localhost:27017

# Atlas
mongosh "mongodb+srv://cluster0.xxxxx.mongodb.net/" --username youruser
```

### From the Application

```bash
# Run the IVR system
go run main.go

# You should see:
# Starting IVR Calling System on port 8080
```

If there's a connection error, you'll see:
```
Failed to initialize database: failed to connect to MongoDB...
```

## Database Structure

The application automatically creates these collections:

1. **campaigns** - Marketing campaigns
2. **calls** - Individual call records
3. **call_logs** - Detailed call event logs

### Indexes Created Automatically

The application creates these indexes for optimal performance:

**campaigns:**
- `name`
- `is_active`

**calls:**
- `campaign_id`
- `status`
- `twilio_call_sid`
- `phone_number`

**call_logs:**
- `call_id`
- `created_at` (descending)

## Viewing Data

### Using mongosh

```bash
mongosh

use ivr_calling_system

# View campaigns
db.campaigns.find().pretty()

# View calls
db.calls.find().pretty()

# View call logs
db.call_logs.find().pretty()

# Count documents
db.campaigns.countDocuments()
db.calls.countDocuments()

# Find specific campaign
db.campaigns.findOne({name: "Test Campaign"})
```

### Using MongoDB Compass (GUI)

1. **Download:**
   - [mongodb.com/products/compass](https://www.mongodb.com/products/compass)

2. **Connect:**
   - Paste your MongoDB URI
   - Click "Connect"

3. **Browse:**
   - Select `ivr_calling_system` database
   - Explore collections visually

## Backup & Restore

### Backup

```bash
# Local MongoDB
mongodump --db ivr_calling_system --out ./backup

# With authentication
mongodump --uri="mongodb+srv://user:pass@cluster.net/" --db ivr_calling_system --out ./backup
```

### Restore

```bash
# Local MongoDB
mongorestore --db ivr_calling_system ./backup/ivr_calling_system

# With authentication
mongorestore --uri="mongodb+srv://user:pass@cluster.net/" --db ivr_calling_system ./backup/ivr_calling_system
```

## Performance Tuning

### For High Volume

1. **Enable Sharding** (Atlas M10+)
2. **Add Read Replicas**
3. **Optimize Queries:**
```javascript
// Add compound indexes
db.calls.createIndex({campaign_id: 1, status: 1})
db.calls.createIndex({created_at: -1, status: 1})
```

### Connection Pooling

The Go driver automatically manages connection pooling. For high-load:

Update `database/database.go`:
```go
clientOptions := options.Client().
    ApplyURI(mongoURI).
    SetMaxPoolSize(100).
    SetMinPoolSize(10)
```

## Troubleshooting

### Connection Timeout

```
Error: failed to ping MongoDB: context deadline exceeded
```

**Solutions:**
- Check MongoDB is running
- Verify network connectivity
- Check firewall rules
- Ensure IP is whitelisted (Atlas)

### Authentication Failed

```
Error: authentication failed
```

**Solutions:**
- Verify username/password in URI
- Check database user permissions
- Ensure special characters in password are URL-encoded

### Cannot Connect to Atlas

**Solutions:**
- Add your IP to whitelist
- Check connection string format
- Verify cluster is running
- Test network connectivity

## Security Best Practices

1. **Use Environment Variables:**
   - Never commit `.env` to git
   - Use different credentials for dev/prod

2. **Enable Authentication:**
   ```bash
   # For local MongoDB
   mongod --auth
   ```

3. **Use SSL/TLS:**
   ```env
   MONGODB_URI=mongodb://localhost:27017/?tls=true&tlsCertificateKeyFile=/path/to/cert.pem
   ```

4. **Restrict Network Access:**
   - Whitelist specific IPs only
   - Use VPC peering (Atlas)

5. **Regular Backups:**
   - Enable automated backups (Atlas)
   - Test restore procedures

## Migration from SQLite

If you previously used SQLite and want to migrate:

1. **Export SQLite data to JSON**
2. **Import to MongoDB:**
```bash
mongoimport --db ivr_calling_system --collection campaigns --file campaigns.json --jsonArray
mongoimport --db ivr_calling_system --collection calls --file calls.json --jsonArray
mongoimport --db ivr_calling_system --collection call_logs --file call_logs.json --jsonArray
```

## Additional Resources

- **MongoDB Documentation:** [docs.mongodb.com](https://docs.mongodb.com)
- **MongoDB University:** [university.mongodb.com](https://university.mongodb.com) (Free courses)
- **Go MongoDB Driver:** [mongodb.com/docs/drivers/go](https://www.mongodb.com/docs/drivers/go/current/)
- **MongoDB Atlas:** [mongodb.com/cloud/atlas](https://www.mongodb.com/cloud/atlas)

---

**Quick Start:** For development, use MongoDB Atlas free tier - it's the fastest way to get started with zero local setup!
