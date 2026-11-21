FROM nginx:1.31.9
COPY nginx/default.conf /etc/nginx/conf.d/default.conf
COPY public /usr/share/nginx/html
