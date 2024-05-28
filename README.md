# Banuba AR Cloud CLI for effects

Banuba AR Cloud service enables your app to connect with our servers and access AR effects directly from the cloud. This means that the effects are downloaded on-demand, whenever a user requires them, instead of being pre-packaged within your app. As a result, your app remains lightweight while still offering a wealth of content.

AR Cloud service can be deployed on your favorite CDN solution i.e. AWS.

## Process data
Use `.zip` archive with AR effects provided by Banuba representatives to process data before uploading on your CDN.

CLI tools are in ```bin``` folder.  
Run the program with the required parameters in Terminal.

### Mac
```bash
./mac_arcloud-effects-cli -s <path/to/effects.zip> -u <your_arcloud_domain>
```

### Windows
```
.\win_arcloud-effects-cli.exe -s <path\to\effects.zip> -u <your_arcloud_domain>
```

### Parameters
- `-s, --source`: Path to the effects zip file `required`
- `-d, --destination`: Destination folder to extract the effects (default: effects)
- `-u, --api-url`: ARCloud domain `required`

```./mac_arcloud-effects-cli -s /Users/username/Work/arcloud-effects-cli/banuba_effects.zip -u https://resources/banuba/```

As a result ```effects``` folder will be created with processed data.  
```effects``` folder includes
1. ```<effect_name>.zip``` and ```<effect_name>.png``` files for each AR effect of the provided `.zip`.
2. ```api_response``` Manifest file that has json structure for all effects.

Example of ```api_response```
```json
{
  "effects": [
    {
      "URL": "https://resources/banuba/Effect1.zip",
      "Preview": "https://resources/banuba/Effect1.png",
      "ETag": "\"c78e0bed9032f3997cb489036040afc6\""
    },
    {
      "URL": "https://resources/banuba/Effect2.zip",
      "Preview": "https://resources/banuba/Effect2.png",
      "ETag": "\"57c5ff0367335807ffc5298cd23c5d33\""
    },
    ...
  ]
}
```
Fields:
- `"URL"` - path to AR effect `.zip` file.
- `"Preview"` - path to AR effect preview `.png` file.
- `"ETag"` - MD5 hash of the effect.

## Upload to CDN
Use data from ```effects``` folder.

### Manifest file
Upload ```api_response``` file to ```https://<your_arcloud_domain>/v1/effects/api_response```

Manifest file is a JSON file without file extension. Set `Content-Type` as `application/json` before upload file to CDN i.e. AWS S3
bucket.

### Effects
Upload all ```.zip``` and ```.png``` files to ```<your_arcloud_domain>```

- `/` (folder store effects and their preview)
    - `Effect1.zip`
    - `Effect1.png>`
- `v1/`
    - `effects/`
        - `api_response`


