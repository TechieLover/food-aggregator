# Command to create a docker build
sudo docker build -t "food-aggregator" ~/work/src/mbrdi/food-aggregator/

# Command to run a container and mapping to external localhost 
sudo docker run -p 127.0.0.1:3000:3000 --env-file ~/work/src/mbrdi/food-aggregator/dev.env "food-aggregator"

# Access to service is http://localhost:3000/ping

# Command to get internal_ip of container by providing container_id , to get container_id use "sudo docker ps" 
sudo docker inspect --format '{{ .NetworkSettings.IPAddress }}' 'container_id'

# Access to service is http://internal_ip:3000/ping

# create a folder mbrdi "mkdir ~/work/src/mbrdi" then copy food-aggregator folder
# To run a service localally then Go to path "cd ~/work/src/mbrdi/food-aggregator" then run "./build.sh"  

# Open Endpoints wrt to localhost
1. http://localhost:3000/v1/buy-item/:item_name --- need to provide a item name
2. http://localhost:3000/v1/buy-item-qty/:item_name?qty={quantity} --- need to provide a item name and quantity in integer
3. http://localhost:3000/v1/buy-item-qty-price/:item_name?qty={quantity}&price={price} --- need to provide a item name, quantity in integer and price in float64
4. http://localhost:3000/v1/show-summary --- need to provide a item name
5. http://localhost:3000/v1/fast-buy-item/:item_name --- need to provide a item name