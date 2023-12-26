# bumder
A mini-project for creating dating apps backend system, 
similar to *Bumble* and *Tinder*.

> [!WARNING]  
 My apology, this project actually is still under development.

## Functional Requirements
1. Login
   - Use mobile phone 
   - It will use 2FA using OTP-based models
   - **[NEXT]** Use google email SSO
2. Register
   - Register using mobile_phone
   - **[NEXT]** Use google email SSO
3. Profile
   - See our bio and preferences
   - Upload up to 3 images per user 
   - **[NEXT]** Subscription for Premium Account
4. Feeds/Recommendation
   - Matchmaking system
   - Swipe right for likes
   - Swipe left for passes
   - Same profile can't appear twice in a day
5. Premium Account
   - Had special labels on feeds appearances
   - No limitation swipe per day
   - Users can view a list of profiles from users who have liked them.
   - **[NEXT]** Super likes, which means if someone matched to/from Super likes

## How To Run Project
> [!IMPORTANT]  
Make sure you already prepare the ***config.yml*** file.  
You can see the config structure example from ***config.yml.example*** file

### Run the application
```bash
go run cmd/server/main.go
```