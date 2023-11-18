# Getting Started
Tana.moe back-end was built with [PocketBase](https://pocketbase.io/), where users could request and send data with RESTful APIs.

## API endpoint
```
https://tana.moe/api/
```

### Images
For fetching and serving pre-processed images, use our [imagor](https://github.com/cshum/imagor#image-endpoint) image processing endpoint.
```
https://image.tana.moe/
```

::: info
The endpoint requires image path with their respective hash (as in, `/unsafe/` are not allowed).
Custom filter arguments are not authorised.
:::

## Next steps
Discover [Tana.moe API reference](/reference).
