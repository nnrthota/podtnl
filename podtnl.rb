class Podtnl < Formula
  desc "Powerful CLI that makes your pod available to online"
  homepage "https://narendranathreddythota.github.io/podtnl/"
  url "https://github.com/narendranathreddythota/podtnl/archive/1.0.tar.gz"
  sha256 "e94d28fae97ee38ee3f4b3a502df5c510280eb49868aeb2cad40a7266158c0cc"
  depends_on "go" => :build
  def install
    system "go", "build", *std_go_args
  end
  test do
    system "#{bin}/podtnl", "-v"
  end
end
