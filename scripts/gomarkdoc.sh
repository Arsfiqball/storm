# Run this command from root project directory
# sh scripts/gomarkdoc.sh
rm -r docs/packages/

# Internal Package
for dir in internal/*/ ; do
    PKG_NAME=$(echo "$dir" | cut -d '/' -f 2- | cut -d '/' -f -1)
    gomarkdoc --output "docs/packages/$PKG_NAME.md" "./$dir"
done

# Public Packages
for dir in pkg/*/ ; do
    PKG_NAME=$(echo "$dir" | cut -d '/' -f 2- | cut -d '/' -f -1)
    gomarkdoc --output "docs/packages/$PKG_NAME.md" "./$dir"
done
