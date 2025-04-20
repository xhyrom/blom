use zed_extension_api as zed;

struct BlomExtension {
    // ... state
}

impl zed::Extension for BlomExtension {
    fn new() -> Self {
        Self {}
    }
}

zed::register_extension!(BlomExtension);
