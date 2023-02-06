FROM nginx:alpine

COPY nginx.conf /etc/nginx/
COPY images /data/images

EXPOSE 80

ENTRYPOINT ["nginx", "-g", "daemon off;"]