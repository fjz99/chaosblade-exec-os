local uri = ngx.var.uri;
if uri == "/tt"
then
    ngx.say(uri);
    ngx.exit(200);
end
