FROM nginx:1.29.0
COPY nginx/default.conf /etc/nginx/conf.d/default.conf
COPY public /usr/share/nginx/html
