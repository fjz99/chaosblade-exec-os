local uri = ngx.var.uri;
if uri == "/tt" or string.match(uri, "/test.*")
then
    ngx.header["a"] = "b"
    ngx.header["Content-Type"] = "text/plain"
    ngx.say(uri);
    ngx.exit(200);
end
