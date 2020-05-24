require "language/go"

class Podtnl < Formula
  desc "Expose your pod to Online easily from any kubernetes clusters without creating a kubernetes service."
  homepage "https://podtnl.sh"
  url "https://github.com/narendranathreddythota/podtnl/archive/1.0.tar.gz"
  sha256 "8f06c335467622419b643ee4f5df54513256c086162f0a83bd7004deae715b6e"
  revision 1
  head "https://github.com/narendranathreddythota/podtnl.git"
  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    path = buildpath/"src/github.com/narendranathreddythota/podtnl"
    path.install Dir["*"]
    cd path do
      system "go", "build", "-o", "#{bin}/podtnl"
    end
  end

  test do
    assert_match version.to_s, shell_output("#{bin}/podtnl", "-v")
  end

end 