FROM nginx:1.27.2
COPY nginx/default.conf /etc/nginx/conf.d/default.conf
COPY public /usr/share/nginx/html
