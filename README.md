# Emby API Go Client

Auto-generated Go client package for [Emby](https://emby.media/). See the [API Documentation](/api/README.md). 

## Example

Generate an API token in the Jellyfin WebUI at Administration -> Overview -> Advanced -> API Token.

## Development

The `configuration.go` file was adjusted (renamed the struct `ServerConfiguration` to `APIServerConfiguration`) and ignored from further generations to avoid conflicts with the same struct from the `model_server_configuration.go` file. In case the `configuration.go` needs to be regenerated, the struct needs to be renamed again.
