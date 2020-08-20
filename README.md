# 💎 Project Capstone

```
 ______   ________   ______   ______   _________  ______   ___   __    ______      
/_____/\ /_______/\ /_____/\ /_____/\ /________/\/_____/\ /__/\ /__/\ /_____/\     
\:::__\/ \::: _  \ \\:::_ \ \\::::_\/_\__.::.__\/\:::_ \ \\::\_\\  \ \\::::_\/_    
 \:\ \  __\::(_)  \ \\:(_) \ \\:\/___/\  \::\ \   \:\ \ \ \\:. `-\  \ \\:\/___/\   
  \:\ \/_/\\:: __  \ \\: ___\/ \_::._\:\  \::\ \   \:\ \ \ \\:. _    \ \\::___\/_  
   \:\_\ \ \\:.\ \  \ \\ \ \     /____\:\  \::\ \   \:\_\ \ \\. \`-\  \ \\:\____/\ 
    \_____\/ \__\/\__\/ \_\/     \_____\/   \__\/    \_____\/ \__\/ \__\/ \_____\/ 
```                                                                                   
                                                                                
Project demonstrating a simple Go webserver, Jenkins for CI/CD and AWS EKS for deployment demo.

## Run tests

Run the simple unit tests to confirm the build is working

```
make test
```

Check the return code for errors (>0 we have a problem Houston!)

```
echo $?
```

## Setup

Build the docker image

```
make docker
```

## Run the image

```
docker -p 8000:8000 capstone-k8
```

Next hit `http://localhost:8000/` for the demo server.
