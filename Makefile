print-last-tag:
	@latest_tag=$$(git describe --tags `git rev-list --tags --max-count=1`); \
	echo "Latest tag is $$latest_tag";

bump-version: print-last-tag
	@echo "Enter the new version (e.g., v1.0.1):"
	@read new_version; \
	echo "Tagging and pushing $$new_version"; \
	git tag $$new_version; \
	git push origin $$new_version; \
	GOPROXY=proxy.golang.org go list -m github.com/papaya147/randomize@$$new_version;

.PHONY: bump-version print-last-tag