def _impl(ctx):
    print("called.")


printer = rule(
    implementation = _impl,
)