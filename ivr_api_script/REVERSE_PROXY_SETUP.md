# 🌐 Reverse Proxy Setup Guide

This guide covers setting up Nginx or Caddy as a reverse proxy for your IVR API at `ivr.waygosquad.com`.

## 📋 Table of Contents

- [Nginx Setup](#nginx-setup)
- [Caddy Setup](#caddy-setup-recommended)
- [SSL Certificates](#ssl-certificates)
- [DNS Configuration](#dns-configuration)
- [Testing](#testing)
- [Troubleshooting](#troubleshooting)

---

## 🔧 Nginx Setup

### Prerequisites

```bash
# Ubuntu/Debian
sudo apt update
sudo apt install nginx certbot python3-certbot-nginx

# CentOS/RHEL
sudo yum install nginx certbot python3-certbot-nginx
```

### Installation Steps

1. **Copy Nginx configuration**:

   ```bash
   sudo cp nginx.conf /etc/nginx/sites-available/ivr.waygosquad.com
   sudo ln -s /etc/nginx/sites-available/ivr.waygosquad.com /etc/nginx/sites-enabled/
   ```

2. **Test configuration**:

   ```bash
   sudo nginx -t
   ```

3. **Obtain SSL certificate**:

   ```bash
   sudo certbot --nginx -d ivr.waygosquad.com
   ```

4. **Start Nginx**:

   ```bash
   sudo systemctl enable nginx
   sudo systemctl start nginx
   ```

5. **Auto-renewal setup**:
   ```bash
   sudo certbot renew --dry-run
   # Certbot automatically sets up a cron job for renewal
   ```

### Nginx Configuration Details

- **HTTP → HTTPS redirect**: Automatic on port 80
- **TLS versions**: TLSv1.2, TLSv1.3
- **Security headers**: HSTS, X-Frame-Options, CSP
- **Proxy timeout**: 60 seconds
- **Max upload size**: 10MB

### Nginx Commands

```bash
# Test configuration
sudo nginx -t

# Reload (no downtime)
sudo nginx -s reload

# Restart
sudo systemctl restart nginx

# View logs
sudo tail -f /var/log/nginx/ivr.waygosquad.com.access.log
sudo tail -f /var/log/nginx/ivr.waygosquad.com.error.log

# Check status
sudo systemctl status nginx
```

---

## 🚀 Caddy Setup (Recommended)

### Why Caddy?

✅ **Automatic HTTPS** - Zero configuration SSL  
✅ **Auto-renewal** - Certificates renew automatically  
✅ **Simpler config** - Easy to read and maintain  
✅ **HTTP/2 & HTTP/3** - Built-in support  
✅ **Zero downtime** - Graceful reloads

### Prerequisites

```bash
# Ubuntu/Debian
sudo apt install -y debian-keyring debian-archive-keyring apt-transport-https
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
sudo apt update
sudo apt install caddy

# Or download binary
# https://caddyserver.com/download
```

### Installation Steps

1. **Copy Caddyfile**:

   ```bash
   sudo cp Caddyfile /etc/caddy/Caddyfile
   ```

2. **Create log directory**:

   ```bash
   sudo mkdir -p /var/log/caddy
   sudo chown caddy:caddy /var/log/caddy
   ```

3. **Test configuration**:

   ```bash
   sudo caddy validate --config /etc/caddy/Caddyfile
   ```

4. **Start Caddy**:
   ```bash
   sudo systemctl enable caddy
   sudo systemctl start caddy
   ```

That's it! Caddy will automatically:

- Obtain SSL certificate from Let's Encrypt
- Configure HTTPS with best practices
- Set up HTTP → HTTPS redirect
- Renew certificates before expiry

### Caddy Configuration Details

- **Automatic HTTPS**: Let's Encrypt integration
- **Auto-renewal**: Certificates renew 30 days before expiry
- **HTTP/2**: Enabled by default
- **Compression**: Gzip and Zstandard
- **Health checks**: 10-second intervals

### Caddy Commands

```bash
# Validate configuration
sudo caddy validate --config /etc/caddy/Caddyfile

# Reload (zero downtime)
sudo caddy reload --config /etc/caddy/Caddyfile

# Restart
sudo systemctl restart caddy

# View logs
sudo journalctl -u caddy -f
sudo tail -f /var/log/caddy/ivr.waygosquad.com.log

# Check status
sudo systemctl status caddy

# Format Caddyfile
caddy fmt --overwrite /etc/caddy/Caddyfile
```

---

## 🔐 SSL Certificates

### Let's Encrypt (Free)

**Nginx**:

```bash
sudo certbot --nginx -d ivr.waygosquad.com
```

**Caddy**: Automatic - no action needed!

### Custom SSL Certificate

**Nginx** - Edit `nginx.conf`:

```nginx
ssl_certificate /path/to/your/certificate.crt;
ssl_certificate_key /path/to/your/private.key;
```

**Caddy** - Edit `Caddyfile`:

```caddy
ivr.waygosquad.com {
    tls /path/to/certificate.crt /path/to/private.key
    # ... rest of config
}
```

---

## 🌍 DNS Configuration

Point your domain to your server:

### A Record

```
Type: A
Name: ivr
Value: YOUR_SERVER_IP
TTL: 3600
```

### Verify DNS propagation:

```bash
# Check DNS
nslookup ivr.waygosquad.com
dig ivr.waygosquad.com

# Wait for propagation (can take 1-48 hours)
```

---

## 🧪 Testing

### 1. Test DNS Resolution

```bash
ping ivr.waygosquad.com
```

### 2. Test HTTP → HTTPS Redirect

```bash
curl -I http://ivr.waygosquad.com
# Should return: 301 or 302 redirect to HTTPS
```

### 3. Test HTTPS

```bash
curl https://ivr.waygosquad.com/health
# Should return: {"status":"healthy"}
```

### 4. Test API Endpoints

```bash
# Health check
curl https://ivr.waygosquad.com/health

# IVR config
curl https://ivr.waygosquad.com/api/v1/config/ivr

# TwiML webhook
curl https://ivr.waygosquad.com/twiml/welcome
```

### 5. Test SSL Certificate

```bash
# Check SSL details
openssl s_client -connect ivr.waygosquad.com:443 -servername ivr.waygosquad.com

# Check SSL expiry
echo | openssl s_client -connect ivr.waygosquad.com:443 2>/dev/null | openssl x509 -noout -dates
```

### 6. Test from Twilio

```bash
# Update your .env file
SERVER_BASE_URL=https://ivr.waygosquad.com

# Restart your Go API
./bin/ivr-api.exe

# Make a test call - Twilio will use HTTPS webhooks
```

---

## 🔧 Troubleshooting

### Nginx Issues

**502 Bad Gateway**:

```bash
# Check if Go API is running
curl http://localhost:8080/health

# Check Nginx error log
sudo tail -f /var/log/nginx/error.log

# Verify proxy_pass URL in nginx.conf
```

**SSL Certificate Error**:

```bash
# Renew certificate
sudo certbot renew

# Check certificate status
sudo certbot certificates
```

**Port 80/443 Already in Use**:

```bash
# Find process using port
sudo lsof -i :80
sudo lsof -i :443

# Stop conflicting service
sudo systemctl stop apache2  # if Apache is running
```

### Caddy Issues

**Failed to bind to port**:

```bash
# Check if another process is using port 80/443
sudo lsof -i :80
sudo lsof -i :443

# Stop Caddy and restart
sudo systemctl stop caddy
sudo systemctl start caddy
```

**Certificate acquisition failed**:

```bash
# Check DNS is pointing to your server
dig ivr.waygosquad.com

# Check firewall allows port 80 and 443
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp

# Check Caddy logs
sudo journalctl -u caddy -n 100
```

**Reverse proxy connection refused**:

```bash
# Verify Go API is running
curl http://localhost:8080/health

# Check Caddy config
sudo caddy validate --config /etc/caddy/Caddyfile

# View logs
sudo journalctl -u caddy -f
```

### Firewall Configuration

```bash
# UFW (Ubuntu)
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw enable

# Firewalld (CentOS/RHEL)
sudo firewall-cmd --permanent --add-service=http
sudo firewall-cmd --permanent --add-service=https
sudo firewall-cmd --reload

# Check if ports are open
sudo netstat -tlnp | grep :80
sudo netstat -tlnp | grep :443
```

### General Debugging

```bash
# Check if API is accessible locally
curl http://localhost:8080/health

# Check reverse proxy logs
# Nginx:
sudo tail -f /var/log/nginx/error.log
# Caddy:
sudo journalctl -u caddy -f

# Test with verbose output
curl -v https://ivr.waygosquad.com/health

# Check TLS handshake
openssl s_client -connect ivr.waygosquad.com:443 -tls1_2
```

---

## 📊 Monitoring

### Nginx

```bash
# Enable status page (add to nginx.conf)
location /nginx_status {
    stub_status;
    allow 127.0.0.1;
    deny all;
}

# View status
curl http://localhost/nginx_status
```

### Caddy

```bash
# Caddy has built-in metrics endpoint
# Add to Caddyfile:
# :2019 {
#     metrics /metrics
# }

# View metrics
curl http://localhost:2019/metrics
```

### Log Analysis

```bash
# Nginx - Most requested endpoints
awk '{print $7}' /var/log/nginx/ivr.waygosquad.com.access.log | sort | uniq -c | sort -rn | head -10

# Nginx - Response codes
awk '{print $9}' /var/log/nginx/ivr.waygosquad.com.access.log | sort | uniq -c | sort -rn

# Caddy - View JSON logs
sudo tail -f /var/log/caddy/ivr.waygosquad.com.log | jq '.'
```

---

## 🎯 Production Checklist

- [ ] DNS A record pointing to server IP
- [ ] Firewall allows ports 80 and 443
- [ ] SSL certificate obtained and valid
- [ ] HTTP → HTTPS redirect working
- [ ] Go API running on localhost:8080
- [ ] Health endpoint accessible via HTTPS
- [ ] Twilio webhooks updated to use HTTPS domain
- [ ] `.env` file has `SERVER_BASE_URL=https://ivr.waygosquad.com`
- [ ] Logs rotating and monitored
- [ ] Auto-renewal configured for SSL

---

## 🚀 Quick Start

**For Caddy (Easiest)**:

```bash
# 1. Install Caddy
sudo apt install caddy

# 2. Copy config
sudo cp Caddyfile /etc/caddy/Caddyfile

# 3. Start
sudo systemctl restart caddy

# Done! ✅
```

**For Nginx**:

```bash
# 1. Install Nginx
sudo apt install nginx certbot python3-certbot-nginx

# 2. Copy config
sudo cp nginx.conf /etc/nginx/sites-available/ivr.waygosquad.com
sudo ln -s /etc/nginx/sites-available/ivr.waygosquad.com /etc/nginx/sites-enabled/

# 3. Get SSL
sudo certbot --nginx -d ivr.waygosquad.com

# 4. Start
sudo systemctl restart nginx

# Done! ✅
```

---

## 📞 Support

If you encounter issues:

1. Check the troubleshooting section above
2. Review logs (Nginx/Caddy and Go API logs)
3. Verify DNS and firewall configuration
4. Test locally first: `curl http://localhost:8080/health`

Your IVR API will be accessible at: **https://ivr.waygosquad.com** 🎉
