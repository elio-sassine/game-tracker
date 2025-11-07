FROM node:alpine AS build

WORKDIR /usr/src/app

COPY package.json package-lock.json ./

RUN npm ci

COPY . ./

RUN npm run build --prod

FROM nginx:alpine

COPY --from=build /usr/src/app/dist/app/browser /usr/share/nginx/html

COPY nginx.conf /etc/nginx/conf.d/default.conf
RUN mkdir -p /certs

# Generate self-signed certs for dev if none are mounted at runtime.
# These will be overwritten by a host mount if you provide real certs.
RUN apk add --no-cache openssl && \
		openssl req -x509 -nodes -days 365 -subj "/CN=localhost" -newkey rsa:2048 \
			-keyout /certs/localhost-key.pem -out /certs/localhost.pem || true

EXPOSE 80 443

CMD ["nginx", "-g", "daemon off;"]