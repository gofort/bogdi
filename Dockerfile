FROM nginx:1.19.7
COPY nginx/default.conf /etc/nginx/conf.d/default.conf
COPY public /usr/share/nginx/html
